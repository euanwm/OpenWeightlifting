"""Main handler file"""
from sport80 import SportEighty

from static_helpers import load_json


class HandlerMain:
    """ This will either update or create new databases """

    def __int__(self, config_file: str):
        self.url_config: dict = load_json(config_file)
        # {'sport80_urls': ['https://bwl.sport80.com/', 'https://usaweightlifting.sport80.com/']}

    def create_index(self) -> dict:
        """ Initial creation of the index file """

    def create_results(self) -> dict:
        """ Initial creation of the results file """

    def update_index(self) -> dict:
        """ Checks that the events index is up-to-date """

    def update_results(self) -> dict:
        """ Adds new results to the end of the database """
