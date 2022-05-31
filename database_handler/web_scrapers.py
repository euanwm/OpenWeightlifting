"""Non-sport80 scraping APIs"""
import requests
from urllib.parse import urljoin


class AustraliaWeightlifting:
    """API class for AWF"""

    AWF_ROOT = "https://www.awf.com.au/"
    INDEX = "competition/statistics/competitions"
    EVENT_PAGE = "competition/statistics/competitions/results/id/"
    HEADERS = {"user-agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) "
                             "Chrome/102.0.0.0 Safari/537.36"}

    def __init__(self):
        self.getter = requests.Session()

    def get_event(self, event_id: int = 2792):
        """Gets the event page - will need to scrape tables etc"""
        # Below link would give you a full PDF results file
        # https://old.awf.com.au/resultsrankings/resultsbook.aspx?compid=2792
        endpoint = urljoin(self.AWF_ROOT, self.EVENT_PAGE + str(event_id))
        index_raw = self.getter.post(endpoint, headers=self.HEADERS)
        print(index_raw.text)


if __name__ == '__main__':
    AWF = AustraliaWeightlifting()
    AWF.get_event()