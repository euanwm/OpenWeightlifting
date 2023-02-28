# https://j-w-a.or.jp/pastrecords/2022/
# Scrape all .pdf links from the page and store them in a list

import requests
import re
import logging as log
from urllib.parse import urljoin
from os import mkdir

log.basicConfig(level=log.INFO)

class Japan:
    def __init__(self):
        self.url = "https://j-w-a.or.jp/pastrecords/"

    def get_pdf_links(self, year: int) -> list:
        """
        Get all pdf links from the page
        :return: list of pdf links
        """
        url = urljoin(self.url, str(year))
        log.info(f"Getting pdf links from {url}")
        page = requests.get(url)
        pdf_links = re.findall(r"https://j-w-a.or.jp/wp/wp-content/uploads/\d{4}/\d{2}/\d{4}_\w+_\w+.pdf", page.text)
        log.info(f"Found {len(pdf_links)} pdfs")
        return pdf_links

    @staticmethod
    def download_pdfs(pdf_links: list) -> None:
        """
        Download all pdfs from the page and store them in a folder named "japan_pdf
        :return: None
        """
        log.info(f"Downloading {len(pdf_links)} pdfs")
        for link in pdf_links:
            pdf_name = link.split("/")[-1]
            pdf = requests.get(link)
            try:
                with open(f"japan_pdf/{pdf_name}", "wb") as file:
                    file.write(pdf.content)
            except FileNotFoundError:
                log.error("Folder japan_pdf not found, creating it now")
                mkdir("japan_pdf")
                with open(f"japan_pdf/{pdf_name}", "wb") as file:
                    file.write(pdf.content)

    def download_year(self, year: int) -> None:
        """
        Download all pdfs from a given year
        :param year: year to download
        :return: None
        """
        log.info(f"Downloading pdfs from {year}")
        pdf_links = self.get_pdf_links(year)
        self.download_pdfs(pdf_links)

if __name__ == '__main__':
    japan = Japan()
    japan.download_year(2022)