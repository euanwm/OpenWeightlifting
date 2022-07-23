"""Non-sport80 scraping APIs"""
from typing import Union

import requests
from urllib.parse import urljoin
from bs4 import BeautifulSoup
from re import search
from static_helpers import write_to_csv


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


def backdate_results(start_id, final_id):
    """meh"""
    awf = AustraliaWeightlifting()
    root_dir = "database_root/AUS"
    for id_int in range(start_id, final_id):
        try:
            write_to_csv(root_dir, id_int, awf.get_event(id_int))
        except AttributeError:
            print(f"no result under event: {id_int}..")


if __name__ == '__main__':
    backdate_results(2798, 3000)
