"""All API calls will go in here to keep things neat"""
import os
from csv import reader
from os.path import join
from typing import Union, Dict

from database_handler import QueryThis, DatabaseEntry
from database_handler.static_helpers import load_csv_as_list, results_to_dict


class GoRESTYourself:
    """Not all pee-pee times are poo-poo times"""

    def __init__(self):
        self.query_root = QueryThis.query_folder
        self.lifter_index = load_csv_as_list(join(self.query_root, "lifter_index.csv"))
        self.db_root = QueryThis.results_root  # meh
        self.leaderboard_data_male = self.__load_file_data('male')
        self.leaderboard_data_female = self.__load_file_data('female')
        self.leaderboard_dict = {'male': self.leaderboard_data_male, 'female': self.leaderboard_data_female}

    def __load_file_data(self, gender: str) -> list[list]:
        """Loads queries data - only for the total leaderboard"""
        if gender not in ('male', 'female'):
            raise Exception("Not a valid gender")
        query_cache_file: str = f"top_total_{gender}.csv"
        with open(join(self.query_root, query_cache_file), 'r', encoding="utf-8") as query_file:
            csv_reader = reader(query_file)
            file_data = [x for x in csv_reader]
        return file_data

    def lifter_totals(self, gender="male", start=0, stop=99) -> Union[list[Dict], Dict]:
        """Default endpoint for the landing page"""
        if gender not in ('male', 'female'):
            return {"error": "gender not valid"}
        dicty_boi: list = []
        try:
            for index, line in enumerate(self.leaderboard_dict[gender][start:stop:]):
                line_struct = DatabaseEntry(*line).__dict__
                line_struct['id'] = index + start
                dicty_boi.append(line_struct)
            return dicty_boi
        except FileNotFoundError:
            return {"get": "fucked"}

    def lifter_sinclairs(self, gender, start, stop):
        """fuck up the shit above"""

    def lifter_lookup(self, lifter_deets: dict):
        """this is gonna be hell to cache"""
        db_path = join(self.db_root, lifter_deets['country'])
        db_csv_paths = os.listdir(db_path)
        all_csv_data = []
        lifter_data = []
        for csv_path in db_csv_paths:
            all_csv_data.append(load_csv_as_list(join(db_path, csv_path)))
        for comps in all_csv_data:
            lifter_data.append([x for x in comps if x[3].lower() == lifter_deets['name']])
        lifter_data = [x[0] for x in lifter_data if len(x) != 0]
        lifter_deets['data'] = results_to_dict(lifter_data)
        return lifter_deets

    def lifter_suggest(self, name: str) -> list[dict]:
        """return a list[dict] of lifter names"""
        search_results = []
        for indx_name, gender, country in self.lifter_index:
            boily_boi = {"name": None, "gender": None, "country": None}
            if name in indx_name:
                boily_boi["name"], boily_boi['gender'], boily_boi["country"] = indx_name, gender, country
                search_results.append(boily_boi)
        return search_results


if __name__ == '__main__':
    api = GoRESTYourself()
    #res = api.lifter_totals()
    #print(res)
    print(api.lifter_totals(gender='female', start=300, stop=303))
    #api.lifter_lookup({'name': 'euan meston', 'gender': 'male', 'country': 'UK'})
