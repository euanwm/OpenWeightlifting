import re

from requests import Session
from bs4 import BeautifulSoup
from datetime import datetime

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




class FranceWeightlifting(WebScraper):
    STARTING_SEASON = 3  # Seasons run from roughly march to march, so 3 is 2019-2020
    LATEST_SEASON = 7
    BASE_URL = "http://scoresheet.ffhaltero.fr/scoresheet/"
    def __init__(self):
        self.session = Session()

    def get_data_by_id(self, id_number):
        pass

    def list_recent_events(self):
        page = self.session.get(f'{self.BASE_URL}{self.LATEST_SEASON}')
        unformatted_table = self.__fetch_main_table(page)
        formatted_table = self.__process_table(unformatted_table)

    def __process_table(self, table):
        # Headings
        # Link, Event Name, Region, Male/Female, Team/Individual, Date, Open/Closed, NAT/INT
        headings = ["Link", "Event Name", "Region", "M/F", "Team/Individual", "Date", "Open/Closed", "NAT/INT"]
        rows = table.find_all('tr')
        # remove the rows that contain "Ouverte"
        rows = [row for row in rows if "Ouverte" not in row.text]
        # do some savvy regex in here later
        for row in rows:
            self.__process_row(row)

    def __process_row(self, row):
        if row.find('a', href=True):
            link = row.find_all('a', href=True)[0]['href']
        if row.find('td'):
            event_name = row.find_all('td')[0].text.strip("\n")
            region = self.__regex_clean(row.find_all('td')[1].text)
            mf = self.__find_and_return(row.find_all('td')[2].text)
            team_ind = self.__find_and_return(row.find_all('td')[3].text)
            date = self.__process_date(row.find_all('td')[4].text.strip("\n"))
        return

    def __regex_clean(self, row) -> str:
        # find the string. start with two spaces and ends with a newline. there are characters and special characters in between
        reggie = re.compile(r"\s\w(.+)\w\n")
        match = reggie.search(row)
        if match:
            return match.group(1).lstrip(" ")

    def __find_and_return(self, row) -> str:
        for french, english in french_to_english.items():
            if french in row.lower():
                return english

    def __process_date(self, date):
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

    def __fetch_main_table(self, page):
        soup = BeautifulSoup(page.text, 'html.parser')
        tables = soup.find_all('table')
        return tables[0]


if __name__ == '__main__':
    f = FranceWeightlifting()
    f.list_recent_events()
