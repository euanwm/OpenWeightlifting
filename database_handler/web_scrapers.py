"""Non-sport80 scraping APIs"""
from typing import Union

import requests
from urllib.parse import urljoin
from bs4 import BeautifulSoup
from re import search


def pull_tables(page_content) -> list:
    """ Returns a dict with details of all the tables within it """
    # debug("pull_tables called")
    soup_parse = BeautifulSoup(page_content.text, "html.parser")
    table_list: list = []
    formatted_table: list = []
    for tables in soup_parse.find_all("table"):
        table_list.append(tables)
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


class AustraliaWeightlifting:
    """API class for AWF"""

    AWF_ROOT = "https://www.awf.com.au/"
    INDEX = "competition/statistics/competitions"
    EVENT_PAGE = "competition/statistics/competitions/results/id/"
    HEADERS = {"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) "
                             "Chrome/102.0.0.0 Safari/537.36"}

    def __init__(self):
        self.getter = requests.Session()

    def get_event(self, event_id: int = 2792):
        """Gets the event page - will need to scrape tables etc"""
        # Below link would give you a full PDF results file
        # https://old.awf.com.au/resultsrankings/resultsbook.aspx?compid=2792
        endpoint = urljoin(self.AWF_ROOT, self.EVENT_PAGE + str(event_id))
        index_raw = self.getter.post(endpoint, headers=self.HEADERS)
        event_name = comp_name(index_raw)
        comp_date(index_raw)
        # dnn_ctr515_dnnTITLE_titleLabel
        # comp name is a HEAD element
        # tables = pull_tables(index_raw)
        # print(tables)


if __name__ == '__main__':
    AWF = AustraliaWeightlifting()
    AWF.get_event()
