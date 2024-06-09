"""Non-sport80 scraping APIs"""
import collections
import os

from datetime import datetime
from typing import Union
from urllib.parse import urljoin
from re import search

import requests
import bs4.element
from bs4 import BeautifulSoup
from .static_helpers import write_to_csv
from .result_dataclasses import IWFHeaders, Result


def pull_tables(
        page_content, table_id=None) -> Union[list[BeautifulSoup], BeautifulSoup]:
    """ Returns a dict with details of all the tables within it """
    # debug("pull_tables called")
    if table_id is None:
        table_id = {}
    soup_parse = BeautifulSoup(page_content.text, "html.parser")
    table_list: list = []
    for tables in soup_parse.find_all("table", table_id):
        table_list.append(tables)
    if len(table_list) == 1:
        return table_list[0]
    return table_list

# pylint: disable=inconsistent-return-statements
# TODO: Error handling for when no return
def comp_name(page_resp) -> str:
    """shitfuckpiss"""
    soup_soup = BeautifulSoup(page_resp.text, "html.parser")
    header = soup_soup.find_all("span", {"class": "Head"})
    if len(header) == 1:
        return header[0].text
    print("Competition name not pulled")


def comp_date(page_resp) -> str:
    """date of comp"""
    soup_soup = BeautifulSoup(page_resp.text, "html.parser")
    date_section = soup_soup.find("div", {"id": "dnn_ctr515_ModuleContent"})
    comp_name_date_raw = date_section.find_all("h2")[0].text
    reg = search("\\d{2}/\\d{2}/\\d{4}", comp_name_date_raw)
    comp_date_raw = reg.group()
    __comp_date = __switch_date_format(comp_date_raw)
    return __comp_date


def __switch_date_format(date: str) -> str:
    """Matches the format of the other databases"""
    date_split = date.split('/')
    day, month, year = date_split
    return f"{year}-{month}-{day}"


def strip_table_headers(table) -> list:
    """ Strips the table headers """
    headers = []
    for tbl_hdr in table.find("tr").find_all("th"):
        headers.append(tbl_hdr.text.strip())
    return headers


def strip_table_body(table):
    """Given a table, returns all its rows plus headers"""
    rows = []
    for tbl_row in table.find_all("tr"):
        cells = []
        tds = tbl_row.find_all("td")
        if len(tds) == 0:
            ths = tbl_row.find_all("th")
            for tbl_hdr in ths:
                if len((data := tbl_hdr.text.strip())) > 0:
                    cells.append(data)
        else:
            for tbl_dat in tds:
                data = tbl_dat.find_all('i')
                if len(data) == 1:
                    strip_it = str(data[0].text).replace(" ", "-")
                    # fixes issue #278
                    if len(strip_it) == 0:
                        strip_it = "0"
                    cells.append(strip_it)
                else:
                    if len((data := tbl_dat.text.strip())) > 0:
                        lift_attr = tbl_dat.attrs.get("style")
                        if lift_attr is not None and "line-through" in lift_attr:
                            data = f"-{data}"
                            cells.append(data)
                        else:
                            cells.append(data)
        rows.append(cells)
    return rows


def funcy_shit(the_shit) -> list:
    """adds category onto end of result line"""
    new = []
    for result_line in the_shit[1::]:
        result_line.insert(0, the_shit[0][0])
        new.append(result_line)
    return new


def table_to_list(shit) -> list:
    """takes the janky html table data and does some weird shit to it"""
    final = []
    new_shit = shit[1::]
    cats_pos = [x for x, y in enumerate(new_shit) if len(y) == 1]
    for index, elem in enumerate(cats_pos):
        try:
            meh = new_shit[elem:cats_pos[index + 1]:]
            final.append(funcy_shit(meh))
        except IndexError:
            meh_2 = new_shit[elem::]
            final.append(funcy_shit(meh_2))
    flatten_list = [item for sublist in final for item in sublist]
    return flatten_list


def remove_unused_columns(big_list: list[list]) -> list[list]:
    """removes the YOB, state, region from the result line"""
    for result_line in big_list:
        del result_line[4:7:]
        del result_line[-1::]
    return big_list


def assign_gender(big_list: list[list]) -> list[list]:
    """DID YOU JUST ASSUME MY GENDER"""
    for list_item in big_list:
        if "F" in list_item[2]:
            list_item[2] = "female"
        elif "M" in list_item[2]:
            list_item[2] = "male"
    return big_list


def highest_csv_id(db_folder: list) -> int:
    """grabs the highest numbered csv"""
    dropped_type = [int(x[:-4]) for x in db_folder]
    dropped_type.sort(reverse=True)
    return dropped_type[0]


class AustraliaWeightlifting:
    """API class for AWF"""

    AWF_ROOT = "https://www.awf.com.au/"
    INDEX = "competition/statistics/competitions"
    EVENT_PAGE = "competition/statistics/competitions/results/id/"
    HEADERS = {"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) "
                            "AppleWebKit/537.36 (KHTML, like Gecko) "
                            "Chrome/102.0.0.0 Safari/537.36"}

    def __init__(self):
        self.getter = requests.Session()

    def lookup_lifter(self, lifter_id: int):
        """/statistics/lifter/id/8600"""

    def get_event(self, event_id: int):
        """Gets the event page - will need to scrape tables etc"""
        endpoint = urljoin(self.AWF_ROOT, self.EVENT_PAGE + str(event_id))
        index_raw = self.getter.post(endpoint, headers=self.HEADERS)
        event_name = comp_name(index_raw)
        event_date = comp_date(index_raw)
        main_table = pull_tables(
            index_raw, {"id": "dnn_ctr515_View_tbResults"})
        raw_results = strip_table_body(main_table)
        final_results = table_to_list(raw_results)
        for result in final_results:
            result.insert(0, event_date)
            result.insert(0, event_name)
        pop_results = remove_unused_columns(final_results)
        gendered_results = assign_gender(pop_results)
        final_header = ['event', 'date', 'gender', 'lifter', 'body_weight_(kg)', 'snatch_lift_1',
                        'snatch_lift_2', 'snatch_lift_3', 'c&j_lift_1', 'c&j_lift_2', 'c&j_lift_3', 
                        'best_snatch', 'best_c&j', 'total']
        gendered_results.insert(0, final_header)
        return gendered_results

    def update_db(self, step=30):
        """meh"""
        root_dir = "../event_data/AUS"
        dir_contents = os.listdir(root_dir)
        # stops writing over the last/highest csv in the directory
        last_id = highest_csv_id(dir_contents) + 1
        for id_int in range(last_id, last_id + step):
            try:
                event = self.get_event(id_int)
                if len(event) > 1:
                    write_to_csv(root_dir, id_int, self.get_event(id_int))
            except AttributeError:
                print(f"no result under event: {id_int}..")

    def rebuild_db(self):
        """ Rebuilds the database by re-fetching all the events currently in the database"""
        root_dir = "../event_data/AUS"
        dir_contents = os.listdir(root_dir)
        for csv_file in dir_contents:
            event_id = int(csv_file.split(".")[0])
            event_data = self.get_event(event_id)
            write_to_csv(root_dir, event_id, event_data)

    def add_single(self, event_id: int):
        """Adds a single event to the database"""
        root_dir = "../event_data/AUS"
        event_data = self.get_event(event_id)
        write_to_csv(root_dir, event_id, event_data)


class InternationalWF:
    """Scraper for the IWF site"""
    IWF_ROOT_URL = "https://iwf.sport"
    # pylint: disable=line-too-long
    # Line is fine, just long
    EVENT_URLS = [
        "https://iwf.sport/results/results-by-events/?event_type=all&event_age=all&event_nation=all",
        "https://iwf.sport/results/results-by-events/results-by-events-old-bw/?event_type=all&event_age=all&event_nation=all",
    ]

    def __init__(self, db_root_dir: str):
        self.iwf_root_dir = db_root_dir

    def fetch_events_list(self) -> list:
        """Returns all the available event IDs"""
        all_bw_data = []
        for url in self.EVENT_URLS:
            req = requests.get(url, timeout=120)
            content = req.text

            soup = BeautifulSoup(content, 'html.parser').find('div', 'cards')
            event_ids = []
            for event_id in soup.find_all('a', 'card', href=True):
                event_ids.append(
                    int(event_id['href'].replace('?event_id=', '')))

            event_name = []
            for event in soup.find_all('span', 'text'):
                event_name.append(event.get_text())

            event_dates = []
            for date in soup.find_all('div', 'col-md-2 col-4 not__cell__767'):
                event_dates.append(date.get_text().strip())

            event_locations = []
            for country in soup.find_all(
                    'div', 'col-md-3 col-4 not__cell__767'):
                event_locations.append(country.get_text().strip())

            zip_it = list(zip(event_ids, event_name,
                          event_dates, event_locations))
            all_bw_data.extend(zip_it)

        all_bw_data = sorted(all_bw_data, key=lambda x: x[0], reverse=False)
        all_bw_data.insert(0, list(IWFHeaders.__annotations__))

        return all_bw_data

    def get_results(self, event_id: int) -> Union[list[dict], bool]:
        """Fetches competition data using the result_id integer"""
        page_data = self.__load_results_page(event_id)
        success, data = self.__scrape_result_info(page_data)
        if success:
            return data
        return False

    def __load_results_page(self, event_id: int) -> BeautifulSoup:
        """Loads the event page for the competition, new weight cats are 441 and above"""
        target_url = f"{self.IWF_ROOT_URL}/results/results-by-events/?event_id={event_id}"
        if event_id <= 440:  # Go on, be pedantic...
            target_url = f"{self.IWF_ROOT_URL}/results/results-by-events/results-by-events-old-bw/?event_id={event_id}"
        res = requests.get(target_url,
                           headers={"Content-Type": "text/html; charset=UTF-8"},
                           timeout=120)
        html = res.text
        return BeautifulSoup(html, "html.parser")

    @staticmethod
    # pylint: disable=too-many-locals, too-many-branches, too-many-statements
    # This method is failing pretty much every measure of complexity. Probably
    # requires a bigger effort to refactor.
    # TODO: Refactor this method to reduce the level of complexity
    def __scrape_result_info(soup_data):
        """Compiles table data into list[dict] format"""
        result_container = soup_data.find_all(
            "div", {"class": "result__container"})
        bw_and_lifts = tuple(Result.__annotations__)[4::]

        if len(result_container) == 0:
            return False, []

        result = []
        for div_id in result_container:
            if (
                    div_id.get("id") == "men_snatchjerk"
                    or div_id.get("id") == "women_snatchjerk"
            ):

                cards_container = div_id.find_all(
                    "div", {"class": "cards"})
                for cards in cards_container[::3]:
                    card_container = cards.find_all(
                        "div", {"class": "card"})

                    for card in card_container[1:]:
                        data_snatch = {}

                        name = card.find_all("p")[1].text.strip()
                        bodyweight = card.find_all(
                            "p")[4].text.strip().split()[1]
                        snatch1 = card.find_all("p")[6].strong.contents[0]
                        snatch2 = card.find_all("p")[7].strong.contents[0]
                        snatch3 = card.find_all("p")[8].strong.contents[0]
                        snatch = card.find_all("p")[9].strong.contents[1]

                        category = (
                            card.parent.previous_sibling.previous_sibling.previous_sibling.previous_sibling.text.strip()
                        )

                        if name and snatch:
                            data_snatch["lifter_name"] = name
                            data_snatch["bodyweight"] = bodyweight
                            data_snatch["snatch_1"] = snatch1
                            data_snatch["snatch_2"] = snatch2
                            data_snatch["snatch_3"] = snatch3
                            data_snatch["best_snatch"] = snatch
                            data_snatch["category"] = category
                        result.append(data_snatch)

                for cards in cards_container[1::3]:
                    card_container = cards.find_all(
                        "div", {"class": "card"})

                    for card in card_container[1:]:
                        data_cj = {}
                        name = card.find_all("p")[1].text.strip()
                        jerk1 = card.find_all("p")[6].strong.contents[0]
                        jerk2 = card.find_all("p")[7].strong.contents[0]
                        jerk3 = card.find_all("p")[8].strong.contents[0]
                        jerk = card.find_all("p")[9].strong.contents[1]

                        if name and jerk:
                            data_cj["lifter_name"] = name
                            data_cj["cj_1"] = jerk1
                            data_cj["cj_2"] = jerk2
                            data_cj["cj_3"] = jerk3
                            data_cj["best_cj"] = jerk

                        result.append(data_cj)

                for cards in cards_container[2::3]:
                    card_container = cards.find_all(
                        "div", {"class": "card"})

                    for card in card_container[1:]:
                        data_total = {}
                        name = card.find_all("p")[1].text.strip()
                        total = card.find_all("p")[8].strong.contents[1]

                        if name and total:
                            data_total["lifter_name"] = name
                            data_total["total"] = total
                        result.append(data_total)

        merged_result = {}
        for res in result:
            key = res["lifter_name"]
            merged_result.setdefault(key, {}).update(res)

        final_table = list(merged_result.values())
        for line in final_table:
            for _, (key, val) in enumerate(line.items()):
                if isinstance(val, bs4.element.Tag):
                    # Annoyingly, double-digit lifts have a space in them
                    line[key] = f"-{val.string.strip(' ')}"
                elif val == '---' and key in bw_and_lifts:
                    line[key] = 0
        return True, final_table

    def __convert_to_conform(
            self, result_data: list[dict], comp_details: list) -> list[list]:
        """ONE OF US, ONE OF US, ONE OF US"""
        insert_me = {'date': self.__conform_date(comp_details[2]),
                     'event': comp_details[1]}
        for data in result_data:
            data.update(insert_me)
        ordered_results = self.__order_correctly(result_data)
        return ordered_results

    def update_results(self) -> None:
        """Looks at comp index and then updates with values not currently saved"""
        comp_index = self.fetch_events_list()
        result_db_ids = [int(x.split(".")[0])
                         for x in os.listdir(self.iwf_root_dir)]
        for comp_info in comp_index[1::]:
            if comp_info[0] not in result_db_ids:
                comp_results = self.get_results(comp_info[0])
                comp_result_data = self.__convert_to_conform(
                    comp_results, comp_info)
                write_to_csv(self.iwf_root_dir, comp_info[0], comp_result_data)

    @staticmethod
    def __conform_date(old_date: str) -> str:
        new_date = datetime.strptime(old_date, "%b %d, %Y")
        return new_date.strftime("%Y-%m-%d")

    @staticmethod
    def __order_correctly(big_data: list[dict]) -> list[list]:
        """Arranges columns to match the standard layout of the Result dataclass"""
        key_order = list(Result.__annotations__)
        for index, line in enumerate(big_data):
            res = collections.OrderedDict()
            for key in key_order:
                if key in line:
                    res[key] = line.pop(key)
            res.update(line.items())
            big_data[index] = dict(res)
        ordered_data = [list(x.values()) for x in big_data]
        ordered_data.insert(0, key_order)
        return ordered_data

    def rebuild_single_event(self, csv_id: int) -> None:
        """Rebuilds a single event"""
        comp_index = self.fetch_events_list()
        for comp_info in comp_index:
            if comp_info[0] == csv_id:
                comp_results = self.get_results(comp_info[0])
                comp_result_data = self.__convert_to_conform(
                    comp_results, comp_info)
                write_to_csv(self.iwf_root_dir, comp_info[0], comp_result_data)
