import json
from os import listdir
from os.path import join

from check_db import EVENT_DATA_PATH, assign_dataclass, load_result_csv_as_list

AUTOBUILD_FILTERS_PATH = "../frontend/autobuild/filter_years.json"

def build_filters() -> None:
    fed_dir = [fed for fed in listdir(EVENT_DATA_PATH) if "." not in fed]
    years_available: list[int] = []
    for fed in fed_dir:
        events = [event for event in listdir(join(EVENT_DATA_PATH, fed)) if "csv" in event]
        for event in events:
            event_path = join(EVENT_DATA_PATH, fed, event)
            event_year = __get_event_year(event_path)
            if event_year not in years_available and event_year != 0:
                years_available.append(event_year)
    __update_filters_json(years_available)

def load_json(filepath: str) -> dict:
    with open(filepath, 'r', encoding='utf-8') as file:
        json_data: dict = json.load(file)
    return json_data

def __get_event_year(event_result_name: str) -> int:
    single_event = load_result_csv_as_list(join(EVENT_DATA_PATH, event_result_name))
    if len(single_event) == 0:
        return 0
    event_date = assign_dataclass(single_event[0]).date
    event_year = int(event_date[:4])
    return event_year

def __update_filters_json(years: list[int]) -> None:
    filters_json = load_json(AUTOBUILD_FILTERS_PATH)
    years = sorted(years, reverse=True)
    for year in years:
        if year not in filters_json:
            year_entry = { str(year): str(year)}
            filters_json["years"].append(year_entry)
    with open(AUTOBUILD_FILTERS_PATH, "w") as f:
        json.dump(filters_json, f, indent=4)

if __name__ == '__main__':
    build_filters()