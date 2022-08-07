"""Main handler file"""
import os

from sport80 import SportEighty
from os.path import join
from csv import writer


class DBHandler:
    """ This will either update or create new databases """

    def __init__(self, url: str, abs_dir: str):
        self.url = url
        self.base_dir = abs_dir
        self.sport80_handler = SportEighty(self.url, return_dict=False)

    def create_results(self, year: int):
        """Yep"""
        # new_funcs = SportEighty(self.url, return_dict=False)
        print(f"creating results for {year}")
        e_index = self.sport80_handler.event_index(year)

        if len(e_index) > 0:
            for _, event_dict in e_index.items():
                self.__write_result_file(event_dict)
            return True
        else:
            return False

    def __write_result_file(self, data_dict: dict):
        """Makes the individual results file"""
        filename = data_dict['action'][0]['route'].split('/')[-1::][0]
        print(f"\tcreating {filename}.csv")
        with open(join(self.base_dir, filename + ".csv"), 'w', encoding="utf-8") as results:
            csv_write = writer(results)
            csv_write.writerows(self.sport80_handler.event_results(data_dict))

    def update_results(self, year: int = 2022):
        """Yep"""
        print(f"updating {self.base_dir.split('/')[1]} database...")
        current_dir = os.listdir(self.base_dir)
        new_funcs = SportEighty(self.url, return_dict=False)
        e_index = new_funcs.event_index(year)
        counter = 0
        for _, event_id in e_index.items():
            if f"{self.__strip_id(event_id['action'][0]['route'])}.csv" not in current_dir:
                self.__write_result_file(event_id)
                counter += 1
        print(f"{counter} file(s) were added")

    @staticmethod
    def __strip_id(event_str: str) -> str:
        """lazy af"""
        return event_str.split('/')[-1::][0]

    @staticmethod
    def __collate_event_id(self, big_dict) -> list:
        """Meh"""
        ids: list = []
        for _, y in big_dict.items():
            ids.append(y['action'][0]['route'].split('/')[-1::][0])
        return ids

    def back_dating(self):
        """Keep going back until there's no results left"""
        start_year = 2022
        while self.create_results(start_year):
            start_year -= 1


if __name__ == '__main__':
    shite = DBHandler("https://usaweightlifting.sport80.com/",
                      "database_root/US")
    # shite.create_results(2022) - uncomment if you're setting it up initially
    shite.update_results()  # comment/delete if you're setting it up initially
    pish = DBHandler("https://bwl.sport80.com/",
                     "database_root/UK")
    # pish.create_results(2022)
    pish.update_results()
