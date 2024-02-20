import os
from abc import ABC, abstractmethod
import logging

from python_tools.database_handler import write_to_csv


class WebScraper(ABC):
    @abstractmethod
    def get_data_by_id(self, id_number):
        pass

    @abstractmethod
    def list_recent_events(self):
        pass


class DBInterface(ABC):
    RESULTS_ROOT = "../backend/event_data/"

    @abstractmethod
    def update_results(self) -> None:
        pass

    @abstractmethod
    def get_single_event(self, event_id):
        pass

    @abstractmethod
    def get_event_list(self):
        logging.info("Getting event list")
        pass

    @abstractmethod
    def generate_result(self, result, eventdata):
        pass
