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

    def create_results(self, year: int = 2022):
        """Yep"""
        new_funcs = SportEighty(self.url, return_dict=False)
        e_index = new_funcs.event_index(year)

        for _, y in e_index.items():
            filename = y['action'][0]['route'].split('/')[-1::][0]
            print(filename)
            with open(join(self.base_dir, filename + ".csv"), 'w', encoding="utf-8") as results:
                csv_write = writer(results)
                csv_write.writerows(new_funcs.event_results(y))

    def check_results(self, year: int = 2022):
        """Yep"""
        print(os.listdir(self.base_dir))
        new_funcs = SportEighty(self.url, return_dict=False)
        e_index = new_funcs.event_index(year)
        ids_only = self.__collate_event_id(e_index)

    def __collate_event_id(self, big_dict):
        """Meh"""
        ids: list = []
        for _, y in big_dict.items():
            ids.append(y['action'][0]['route'].split('/')[-1::][0])


if __name__ == '__main__':
    shite = DBHandler("https://usaweightlifting.sport80.com/",
                      "/database_root/US")
    # shite.create_results()
    # shite.check_results()
