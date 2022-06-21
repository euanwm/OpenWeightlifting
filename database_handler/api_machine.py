"""All API calls will go in here to keep things neat"""
import json
from csv import reader
from os.path import join
from pprint import pprint
from typing import Union, Dict

from query_machine import QueryThis
from result_dataclasses import DatabaseEntry


class GoRESTYourself:
    """Not all pee-pee times are poo-poo times"""

    def __init__(self):
        self.query_root = QueryThis.query_folder

    def lifter_totals(self, gender="male", start=0, stop=100) -> Union[dict[str, str], str]:
        """Default endpoint for the landing page"""
        query_filename: str = f"top_total_{gender}.csv"
        file_data: list = []
        dicty_boi:dict = {}
        try:
            with open(join(self.query_root, query_filename), 'r', encoding="utf-8") as query_file:
                csv_reader = reader(query_file)
                file_data = [x for x in csv_reader]
        except FileNotFoundError:
            return {"get": "fucked"}
        finally:
            for index, line in enumerate(file_data[start:stop:]):
                line_struct = DatabaseEntry(*line)
                dicty_boi[str(index + start)] = line_struct.__dict__
            return json.dumps(dicty_boi)


if __name__ == '__main__':
    api = GoRESTYourself()
    res = api.lifter_totals()
    print(res)
