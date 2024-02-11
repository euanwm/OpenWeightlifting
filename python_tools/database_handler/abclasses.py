from abc import ABC, abstractmethod


class WebScraper(ABC):
    @abstractmethod
    def get_data_by_id(self, id_number):
        pass

    @abstractmethod
    def list_recent_events(self):
        pass
