"""All API calls will go in here to keep things neat"""
from csv import reader
from os.path import join

from query_machine import QueryThis
from result_dataclasses import DatabaseEntry


class GoRESTYourself:
    """Not all pee-pee times are poo-poo times"""

    def __init__(self):
        self.query_root = QueryThis.query_folder

    def lifter_totals(self, gender="male", start=0, stop=100) -> dict:
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
                dicty_boi[index + start] = line_struct.__dict__
            return dicty_boi


if __name__ == '__main__':
    api = GoRESTYourself()
    api.lifter_totals()
