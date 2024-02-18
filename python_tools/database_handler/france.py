import re

from requests import Session
from bs4 import BeautifulSoup
from datetime import datetime
from dataclasses import dataclass

from abclasses import WebScraper

french_months = {
    "Jan": 1,
    "Fév": 2,
    "Mar": 3,
    "Avr": 4,
    "Mai": 5,
    "Jui": 6,
    "Jul": 7,
    "Aoû": 8,
    "Sep": 9,
    "Oct": 10,
    "Nov": 11,
    "Déc": 12,
}

french_to_english = {
    "mixte": "Mixed",
    "masculin": "Men",
    "feminin": "Women",
    "equipes": "Teams",
    "individuel": "Individual",

}


@dataclass
class FranceEventInfo:
    link: str
    event_name: str
    region: str
    male_female: str
    team_ind: str
    date: str
    open_closed: str
    nat_int: str


class FranceWeightlifting(WebScraper):
    STARTING_SEASON = 3  # Seasons run from roughly march to march, so 3 is 2019-2020
    LATEST_SEASON = 7
    BASE_URL = "http://scoresheet.ffhaltero.fr/scoresheet/"
    RESULTS_URL = "competition/view/"

    def __init__(self):
        self.session = Session()

    def get_data_by_id(self, id_number):
        page = self.session.get(f'{self.BASE_URL}{self.RESULTS_URL}{id_number}')
        soup = BeautifulSoup(page.text, 'html.parser')
        table = soup.find_all('table')[1]
        results = []
        headers = table.find_all('th')
        header_row = []
        for header in headers:
            header_row.append(header.text)
        rows = table.find_all('tr')
        for row in rows:
            cells = row.find_all('td')
            processed_row = []
            for i, cell in enumerate(cells):
                if i in [0, 2]:
                    processed_row.append(self.__regex_simple_number(cell))
                if i in [1, 3, 4]:
                    processed_row.append(self.__regex_short_clean(cell.text))
                if i in [5, 17]:
                    processed_row.append(self.__regex_float_number(cell))
                if i in [15, 16]:
                    processed_row.append(self.__regex_short_clean(cell.text))
                if i in [6, 7, 8, 9, 10, 11, 12, 13, 14]:
                    processed_row.append(self.__process_score(cell))
            results.append(processed_row)
        return results

    def __regex_float_number(self, cell) -> str:
        reggie = re.compile(r"(\d+,\d+)")
        match = reggie.search(cell.text)
        if match:
            return match.group(1)

    def __regex_simple_number(self, cell) -> str:
        reggie = re.compile(r"\s(\d+)\n")
        match = reggie.search(cell.text)
        if match:
            return match.group(1)
    def __process_score(self, cell) -> str:
        reggie = re.compile(r"\n(-?\d+)\n")
        match = reggie.search(cell.text)
        if match:
            return match.group(1)

    def list_recent_events(self) -> list[FranceEventInfo]:
        page = self.session.get(f'{self.BASE_URL}{self.LATEST_SEASON}')
        unformatted_table = self.__fetch_main_table(page)
        formatted_table = self.__process_table(unformatted_table)
        return formatted_table

    def __process_table(self, table) -> list[FranceEventInfo]:
        rows = table.find_all('tr')
        # remove the rows that contain "Ouverte"
        rows = [row for row in rows if "Ouverte" not in row.text]
        hydrated_table = []
        for row in rows:
            hydrated_table.append(self.__process_row(row))

        # remove any None values
        # for some reason there's a None value first in the list due to the first row of the table being filters
        hydrated_table = [x for x in hydrated_table if x]
        return hydrated_table

    def __process_row(self, row) -> FranceEventInfo:
        processed_data = None
        if row.find('a', href=True) and row.find_all('td'):
            processed_data = FranceEventInfo(
                link=row.find_all('a', href=True)[0]['href'],
                event_name=row.find_all('td')[0].text.strip("\n"),
                region=self.__regex_clean(row.find_all('td')[1].text),
                male_female=self.__find_and_return(row.find_all('td')[2].text),
                team_ind=self.__find_and_return(row.find_all('td')[3].text),
                date=self.__process_date(row.find_all('td')[4].text.strip("\n")),
                open_closed=self.__find_and_return(row.find_all('td')[5].text),
                nat_int=self.__regex_short_clean(row.find_all('td')[6].text)
            )
        return processed_data

    @staticmethod
    def __regex_short_clean(row) -> str:
        reggie = re.compile(r"\s\w(.*)\n")
        match = reggie.search(row)
        if match:
            trimmed = match.group(0).lstrip(" ").rstrip("\n")
            return trimmed

    @staticmethod
    def __regex_clean(row) -> str:
        # find the string. start with two spaces and ends with a newline. there are characters and special characters in between
        reggie = re.compile(r"\s\w(.+)\w\n")
        match = reggie.search(row)
        if match:
            return match.group(1).lstrip(" ")

    @staticmethod
    def __find_and_return(row) -> str:
        for french, english in french_to_english.items():
            if french in row.lower():
                return english

    @staticmethod
    def __process_date(date):
        # 10 Fév 2024
        reggie = re.compile(r"(\d{2}) (\D{3}) (\d{4})")
        match = reggie.search(date)
        if match:
            day = match.group(1)
            month = french_months[match.group(2)]
            year = match.group(3)
            # return formatted date in DD-MM-YYYY
            return datetime(int(year), month, int(day)).strftime("%d-%m-%Y")
        return

    @staticmethod
    def __fetch_main_table(page):
        soup = BeautifulSoup(page.text, 'html.parser')
        tables = soup.find_all('table')
        return tables[0]


if __name__ == '__main__':
    f = FranceWeightlifting()
    # f.list_recent_events()
    f.get_data_by_id(7839)
