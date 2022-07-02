"""All API calls will go in here to keep things neat"""
import json
from csv import reader
from os.path import join
from typing import Union, Dict

from database_handler.query_machine import QueryThis
from database_handler.result_dataclasses import DatabaseEntry


class GoRESTYourself:
    """Not all pee-pee times are poo-poo times"""

    def __init__(self):
        self.query_root = QueryThis.query_folder

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
            return json.dumps(dicty_boi)
        except FileNotFoundError:
            return {"get": "fucked"}

    def lifter_sinclairs(self, gender, start, stop):
        """fuck up the shit above"""

    def lifter_lookup(self, name: str):
        """this is gonna be hell to cache"""
        return {"name": name, "history": ["comp1", "comp2", "comp3"]}

if __name__ == '__main__':
    api = GoRESTYourself()
    res = api.lifter_totals()
    print(res)
