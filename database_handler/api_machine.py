"""All API calls will go in here to keep things neat"""
import csv
import json
from csv import reader
from os.path import join
from typing import Union, Dict

from database_handler.query_machine import QueryThis
from database_handler.result_dataclasses import DatabaseEntry
from database_handler.static_helpers import load_csv_as_list


class GoRESTYourself:
    """Not all pee-pee times are poo-poo times"""

    def __init__(self):
        self.query_root = QueryThis.query_folder
        self.lifter_index = load_csv_as_list(join(self.query_root, "lifter_names.csv"))

    def lifter_totals(self, gender="male", start=0, stop=100) -> Union[dict[str, str], str]:
        """Default endpoint for the landing page"""
        query_cache_file: str = f"top_total_{gender}.csv"
        dicty_boi: dict = {}
        try:
            with open(join(self.query_root, query_cache_file), 'r', encoding="utf-8") as query_file:
                csv_reader = reader(query_file)
                file_data = [x for x in csv_reader]
            for index, line in enumerate(file_data[start:stop:]):
                line_struct = DatabaseEntry(*line)
                dicty_boi[str(index + start)] = line_struct.__dict__
            return dicty_boi
        except FileNotFoundError:
            return {"get": "fucked"}

    def lifter_sinclairs(self, gender, start, stop):
        """fuck up the shit above"""

    def lifter_lookup(self, name: dict):
        """this is gonna be hell to cache"""
        return {"name": name, "history": ["comp1", "comp2", "comp3"]}

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
    #print(api.lifter_suggest("euan"))
    api.lifter_lookup({'name': 'euan warren', 'gender': 'male', 'country': 'AUS'})