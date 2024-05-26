import logging
import os
import re

from requests import Session
from bs4 import BeautifulSoup
from datetime import datetime
from dataclasses import dataclass

from .result_dataclasses import Result
from .static_helpers import results_to_csv, load_json
from .abclasses import WebScraper, DBInterface

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

    def get_event_title(self):
        name = self.event_name.split(" - ")
        if len(name) > 1:
            return name[1]


@dataclass
class FranceResult:
    license: str
    name: str
    birthyear: str
    club: str
    nation: str
    bodyweight: str
    snatch_1: str
    snatch_2: str
    snatch_3: str
    best_snatch: str
    cj_1: str
    cj_2: str
    cj_3: str
    best_cj: str
    total: str
    series: str
    category: str
    iwf_points: str


@dataclass
class FranceEventMetadata:
    jury_1: str
    jury_2: str
    jury_3: str
    referee_1: str
    referee_2: str
    referee_3: str
    timekeeper: str
    technical_controller: str
    marshal: str
    secretary: str
    announcer: str
    trainee_referee_1: str
    trainee_referee_2: str
    trainee_referee_3: str


class FranceWeightlifting(WebScraper):
    STARTING_SEASON = 3  # Seasons run from roughly march to march, so 3 is 2019-2020
    LATEST_SEASON = 7
    BASE_URL = "http://scoresheet.ffhaltero.fr/scoresheet/"
    RESULTS_URL = "competition/view/"
    FEDERATION_SHORTHAND = "FFH"

    def __init__(self):
        self.session = Session()

    def get_data_by_id(self, id_number) -> list[FranceResult] | None:
        page = self.session.get(f'{self.BASE_URL}{self.RESULTS_URL}{id_number}')
        soup = BeautifulSoup(page.text, 'html.parser')
        if len(soup.find_all('table')) < 2:
            return None
        metadata_table = self.__process_metadata(soup.find_all('table')[0])
        tables = soup.find_all('table')[1:]
        results = []
        for table in tables:
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
                if len(processed_row) > 0:
                    results.append(self.__process_result(processed_row))
        return results

    def __process_metadata(self, table) -> type[FranceEventMetadata]:
        # todo: this is a placeholder until I can be bothered to write the code to process the metadata
        return FranceEventMetadata

    def __process_result(self, row) -> FranceResult:
        return FranceResult(
            license=row[0],
            name=row[1],
            birthyear=row[2],
            club=row[3],
            nation=row[4],
            bodyweight=row[5],
            snatch_1=row[6],
            snatch_2=row[7],
            snatch_3=row[8],
            best_snatch=row[9],
            cj_1=row[10],
            cj_2=row[11],
            cj_3=row[12],
            best_cj=row[13],
            total=row[14],
            series=row[15],
            category=row[16],
            iwf_points=row[17]
        )

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

    def list_recent_events(self, season=LATEST_SEASON) -> list[FranceEventInfo] | None:
        page = self.session.get(f'{self.BASE_URL}{season}')
        if not page.ok:
            return None
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
            # return formatted date in YYYY-MM-DD
            return datetime(int(year), month, int(day)).strftime("%Y-%m-%d")
        return

    @staticmethod
    def __fetch_main_table(page):
        soup = BeautifulSoup(page.text, 'html.parser')
        tables = soup.find_all('table')
        return tables[0]


@dataclass
class CollatedEvent:
    date: str
    event_title_short: str
    events: list[FranceEventInfo]


class FranceInterface(DBInterface):
    def __init__(self):
        self.f = FranceWeightlifting()
        self.RESULTS_PATH = os.path.join(self.RESULTS_ROOT, self.f.FEDERATION_SHORTHAND)
        self.CATEGORIES = load_json(
            f"{os.getcwd()}/database_handler/gender_categories.json")
        self.NEXT_SEASON_CHECKED = False
        self.NEXT_AVAILABLE_CSV_ID = 1

    def check_available_csv_id(self) -> int:
        csv_ids = max([0], [int(x.split(".")[0]) for x in os.listdir(self.RESULTS_PATH)])
        self.NEXT_AVAILABLE_CSV_ID = max(csv_ids) + 1
        return self.NEXT_AVAILABLE_CSV_ID

    def get_event_list(self):
        return self.f.list_recent_events()

    def get_single_event(self, event_link):
        event_link = event_link.split("/")[-1]
        return self.f.get_data_by_id(event_link)

    def update_results(self):
        logging.info("Updating results")
        event_list = self.get_event_list()
        number_of_events_added = 0
        result_db_ids = [int(x.split(".")[0])
                         for x in os.listdir(self.RESULTS_PATH)]
        if event_list is not None:
            for event in event_list:
                event_id = int(event.link.split('/')[-1])
                if event_id not in result_db_ids:
                    print(f"Getting results for {event.event_name} / {event.link.split('/')[-1]}")
                    event_results = self.get_single_event(event.link)
                    if event_results is not None:
                        amal_data = []
                        for result in event_results:
                            amal_data.append(self.generate_result(result, event))
                        if amal_data:
                            results_to_csv(self.RESULTS_PATH, event.link.split('/')[-1], amal_data)
                            number_of_events_added += 1
                    else:
                        print(f"No results logged for {event.event_name} / {event.link.split('/')[-1]}")
        if number_of_events_added == 0 and not self.NEXT_SEASON_CHECKED:
            print("No new events added, checking next season")
            self.f.LATEST_SEASON += 1
            self.NEXT_SEASON_CHECKED = True
            self.update_results()

    def generate_result(self, result: FranceResult, eventdata: FranceEventInfo) -> Result | None:
        amal_data = Result(
            event=eventdata.event_name,
            date=eventdata.date,
            category=self.conform_categories(result.category.replace(u'\xa0', ' ')),
            lifter_name=result.name.replace(u'\xa0', ' '),
            bodyweight=float(result.bodyweight.replace(",", ".")),
            snatch_1=float(result.snatch_1),
            snatch_2=float(result.snatch_2),
            snatch_3=float(result.snatch_3),
            cj_1=float(result.cj_1),
            cj_2=float(result.cj_2),
            cj_3=float(result.cj_3),
        )
        return amal_data

    def conform_categories(self, category: str) -> str:
        if " M " in category:
            return category.replace(" M ", " Men ")
        if " F " in category:
            return category.replace(" F ", " Women ")

    def build_old_database(self):
        # this is a one-hitter to build the database from the old naming convention
        # you'll need to run the new_build_database func to collect the rest
        for n in range(3, 5):
            events_list = self.f.list_recent_events(n)
            result_db_ids = [int(x.split(".")[0])
                             for x in os.listdir(self.RESULTS_PATH)]
            for event in events_list:
                event_id = int(event.link.split('/')[-1])
                number_of_events_added = 0
                if event_id not in result_db_ids:
                    print(f"Getting results for {event.event_name} / {event.link.split('/')[-1]}")
                    event_results = self.get_single_event(event.link)
                    if event_results is not None:
                        amal_data = []
                        for result in event_results:
                            amal_data.append(self.generate_result(result, event))
                        if amal_data:
                            results_to_csv(self.RESULTS_PATH, event.link.split('/')[-1], amal_data)
                            number_of_events_added += 1
                    else:
                        print(f"No results logged for {event.event_name} / {event.link.split('/')[-1]}")

    def new_build_database(self):
        starting_csv_id = self.check_available_csv_id()
        # indent and range from here
        for x in range(4, 8):
            events_list = self.f.list_recent_events(x)
            collated_event_info = self.collate_event_ids(events_list)
            print(f"Season {3}: {len(collated_event_info)} colllated event vs {len(events_list)} events")
            for index, event in collated_event_info.items():
                amal_data = []
                for result in event.events:
                    event_results = self.get_single_event(result.link)
                    if event_results is not None:
                        for result in event_results:
                            amal_data.append(self.generate_result(result, event.events[0]))
                if amal_data:
                    results_to_csv(self.RESULTS_PATH, starting_csv_id, amal_data)
                    starting_csv_id += 1
                print(f"Event {index} done")


    def collate_event_ids(self, event_list: list[FranceEventInfo]) -> dict[int, CollatedEvent]:
        event_dict = {}
        index_ticker = self.check_available_csv_id()
        for event in event_list:
            in_event_dict, dict_index = self.iter_dict(event_dict, event)
            if not in_event_dict:
                event_dict[index_ticker] = CollatedEvent(
                    date=event.date,
                    event_title_short=event.get_event_title(),
                    events=[event]
                )
                index_ticker += 1
            else:
                event_dict[dict_index].events.append(event)
        return event_dict

    def iter_dict(self, event_dict: dict[int, CollatedEvent], single_event: FranceEventInfo) -> (bool, int):
        for k, v in event_dict.items():
            if v.date == single_event.date and v.event_title_short == single_event.get_event_title():
                event_dict[k].events.append(single_event)
                return True, k
        return False, None


if __name__ == '__main__':
    # f = FranceWeightlifting()
    # f.list_recent_events()
    f = FranceInterface()
    # this = f.collate_event_ids()
    f.new_build_database()
    pass
    # f.get_data_by_id(7839)
    # f.update_results()
    # f.build_database()
