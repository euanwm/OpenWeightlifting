"""Non-sport80 scraping APIs"""
import os
from _csv import reader
from typing import Union

import requests
from urllib.parse import urljoin
from bs4 import BeautifulSoup
from re import search, sub
from .static_helpers import write_to_csv
from .result_dataclasses import IWFHeaders


def pull_tables(page_content, id=None) -> Union[list[BeautifulSoup], BeautifulSoup]:
    """ Returns a dict with details of all the tables within it """
    # debug("pull_tables called")
    if id is None:
        id = {}
    soup_parse = BeautifulSoup(page_content.text, "html.parser")
    table_list: list = []
    for tables in soup_parse.find_all("table", id):
        table_list.append(tables)
    if len(table_list) == 1:
        return table_list[0]
    return table_list


def comp_name(page_resp) -> str:
    """shitfuckpiss"""
    soup_soup = BeautifulSoup(page_resp.text, "html.parser")
    header = soup_soup.find_all("span", {"class": "Head"})
    if len(header) == 1:
        return header[0].text
    else:
        print("Competition name not pulled")


def comp_date(page_resp) -> str:
    """date of comp"""
    soup_soup = BeautifulSoup(page_resp.text, "html.parser")
    date_section = soup_soup.find("div", {"id": "dnn_ctr515_ModuleContent"})
    comp_name_date_raw = date_section.find_all("h2")[0].text
    reg = search("\d{2}/\d{2}/\d{4}", comp_name_date_raw)
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
                    cells.append(strip_it)
                else:
                    if len((data := tbl_dat.text.strip())) > 0:
                        cells.append(data)
        rows.append(cells)
    return rows


def funcy_shit(the_shit) -> list:
    """adds category onto end of result line"""
    new = []
    for x in the_shit[1::]:
        x.insert(0, the_shit[0][0])
        new.append(x)
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
            meh_2 = (new_shit[elem::])
            final.append(funcy_shit(meh_2))
    flatten_list = [item for sublist in final for item in sublist]
    return flatten_list


def remove_unused_columns(big_list: list[list]) -> list[list]:
    """removes the YOB, state, region from the result line"""
    for x in big_list:
        del x[4:7:]
        del x[-1::]
    return big_list


def assign_gender(big_list: list[list]) -> list[list]:
    """DID YOU JUST ASSUME MY GENDER"""
    for x in big_list:
        if "F" in x[2]:
            x[2] = "female"
        elif "M" in x[2]:
            x[2] = "male"
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
    HEADERS = {"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) "
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
        main_table = pull_tables(index_raw, {"id": "dnn_ctr515_View_tbResults"})
        raw_results = strip_table_body(main_table)
        final_results = table_to_list(raw_results)
        for x in final_results:
            x.insert(0, event_date)
            x.insert(0, event_name)
        pop_results = remove_unused_columns(final_results)
        gendered_results = assign_gender(pop_results)
        final_header = ['event', 'date', 'gender', 'lifter', 'body_weight_(kg)', 'snatch_lift_1', 'snatch_lift_2',
                        'snatch_lift_3', 'c&j_lift_1', 'c&j_lift_2', 'c&j_lift_3', 'best_snatch', 'best_c&j', 'total']
        gendered_results.insert(0, final_header)
        return gendered_results

    def update_db(self, step=30):
        """meh"""
        root_dir = "database_root/AUS"
        print(f"updating {root_dir.split('/')[1]} database...")
        dir_contents = os.listdir(root_dir)
        last_id = highest_csv_id(dir_contents) + 1  # stops writing over the last/highest csv in the directory
        for id_int in range(last_id, last_id + step):
            try:
                write_to_csv(root_dir, id_int, self.get_event(id_int))
            except AttributeError:
                print(f"no result under event: {id_int}..")


class InternationalWF:
    """Scraper for the IWF site"""
    EVENT_URLS = ["https://iwf.sport/results/results-by-events/?event_type=all&event_age=all&event_nation=all",
                  "https://iwf.sport/results/results-by-events/results-by-events-old-bw/?event_type=all&event_age=all&event_nation=all"]

    def __int__(self):
        self.iwf_root_dir = "database_root/IWF"

    def fetch_events(self) -> list:
        """Returns all the available event IDs"""
        all_bw_data = []
        for url in self.EVENT_URLS:
            req = requests.get(url)
            content = req.text

            soup = BeautifulSoup(content, 'html.parser').find('div', 'cards')
            event_ids = []
            for event_id in soup.find_all('a', 'card', href=True):
                event_ids.append(int(event_id['href'].replace('?event_id=', '')))

            event_name = []
            for event in soup.find_all('span', 'text'):
                event_name.append(event.get_text())

            event_dates = []
            for date in soup.find_all('div', 'col-md-2 col-4 not__cell__767'):
                event_dates.append(date.get_text().strip())

            event_locations = []
            for country in soup.find_all('div', 'col-md-3 col-4 not__cell__767'):
                event_locations.append(country.get_text().strip())

            zip_it = list(zip(event_ids, event_name, event_dates, event_locations))
            # The below line could stay in, rest of the code doesn't really care that it's a list of tuples vs a list
            # zip_it = [list(x) for x in zip_it]
            all_bw_data.extend(zip_it)

        all_bw_data = sorted(all_bw_data, key=lambda x: x[0], reverse=False)
        all_bw_data.insert(0, [x for x in IWFHeaders.__annotations__])

        return all_bw_data

    def is_cat(self, line):
        return (("Men" in line) | ("Women" in line)) & ("kg" in line)

    def is_sec(self, line):
        return line in ["Snatch", "Clean&Jerk", "Total"]

    def is_head(self, line, headers):
        return line in headers

    def rep_list(self, x: str, matches: list):
        for match in matches:
            x = x.replace(match, "")
        return x.strip()

    def get_text(self, soup: BeautifulSoup) -> str:
        """Soup in, text out."""
        text = soup.get_text()
        lines = (line.strip() for line in text.splitlines())
        chunks = (phrase.strip() for line in lines for phrase in line.split("  "))
        text = "\n".join(chunk for chunk in chunks if chunk)
        return text

    def containsNumber(self, value) -> bool:
        """Is the value a valid number?"""
        for character in value:
            if character.isdigit():
                return True
        return False

    def scrape_url(self, event_id: int) -> None:
        """Scrapes results page

        Args:
            event_id (int): id of event.
        """

        url = f"https://iwf.sport/results/results-by-events/?event_id={event_id}"
        old_bw_class = False
        headers = ["Rank:", "Name:", "Nation:", "Born:", "B.weight:", "Group:", "1:", "2:", "3:", "Total:", "Snatch:",
                   "CI&Jerk:"]
        #  OK, not a great way to do this but if I refactor it any further it'll be a breaking change
        csv_headers = ["rank", "name", "nation", "born", "bw", "group", "lift1", "lift2", "lift3", "lift4", "cat",
                       "sec",
                       "event_id", "old_classes"]

        if event_id < 441:
            # Changeover of BW categories
            old_bw_class = True
            url = f"https://iwf.sport/results/results-by-events/results-by-events-old-bw/?event_id={event_id}"

        req = requests.get(url)
        content = req.text.replace("<strike>", "-").replace("</strike>", "")
        soup = BeautifulSoup(content, "html.parser")

        event = soup.find("h2").text

        men = self.get_text(soup.find("div", {"id": "men_snatchjerk"})).splitlines()
        women = self.get_text(soup.find("div", {"id": "women_snatchjerk"})).splitlines()

        both_genders = men + women

        row = []
        big_data = []
        for line in both_genders:
            if self.is_cat(line):
                cat = line.replace(" ", "")
            elif self.is_sec(line):
                sec = line.replace("&", "")
            elif self.is_head(line, headers):
                col = 1
            else:
                row.append(self.rep_list(line, headers))
                if "Total:" in line:

                    while len(row) < 10:  # a
                        row.append('---')
                    row.extend([cat, sec, event_id, old_bw_class])
                    big_data.append(row)
                    row = []

        big_data.insert(0, csv_headers)
        filename = f"{event_id}_{self.gen_filename(event)}"
        write_to_csv(f"{dir}/../raw_data/results/", filename, big_data)

        print(f"Event ID: {event_id}")
        print(f"Event: {event}")
        print(f"Saved To: {filename}.csv\n")

    def gen_filename(self, raw_name: str) -> str:
        """Strips all special characters for saving operations"""
        new_name = sub(r"[^a-zA-Z0-9]", "_", raw_name)
        return new_name

    def scrape_pass_errors(self, event_id) -> int:
        """Checks whether an event is present in the results folder and adds it if required"""
        existing_ids = self.fetch_result_ids()

        if event_id in existing_ids:
            return 1
        elif event_id not in existing_ids:
            self.scrape_url(event_id)
            return 0
        else:
            return 2

    def updateResults(self) -> None:
        """Updates the raw data results in-line with the events.csv"""
        event_ids = self.fetch_event_ids()
        results = list(map(self.scrape_pass_errors, event_ids))
        print(results.count(0), "events scraped")
        print(results.count(1), "events already scraped")
        print(results.count(2), "events failed to scrape")

    def fetch_event_ids(self) -> list:
        """Checks the events.csv file in raw_data and returns all event IDs"""
        event_ids: list = []
        with open(f"{dir}/../raw_data/events.csv", "r", encoding='utf-8') as results_file:
            csv_read = reader(results_file)
            for lines in csv_read:
                event_ids.append(lines[0])
        event_ids = [int(x) for x in event_ids[1::]]  # I'm lazy but fuck Pandas/Numpy
        return event_ids

    def fetch_result_ids(self) -> list[int]:
        """Split this down into a function instead of a single line as it's not that Pythonic/readable"""
        result_filenames = os.listdir(self.iwf_root_dir)
        result_ids = [x.split("_")[0] for x in result_filenames]  # todo: this will need changed
        result_ids_as_int = list(map(int, result_ids))
        return result_ids_as_int
