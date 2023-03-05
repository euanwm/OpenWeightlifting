import requests
import re
import logging as log

import tabula

from urllib.parse import urljoin
from os import mkdir, listdir, path
from shutil import move

from PyPDF2 import PdfReader, PageObject
from pandas import DataFrame

log.basicConfig(level=log.INFO)

class Japan:
    """
    Class to scrape pdfs from https://j-w-a.or.jp/pastrecords/
    """
    def __init__(self):
        self.url = "https://j-w-a.or.jp/pastrecords/"
        self.pdf_root = "japan_pdf"

    def get_pdf_links(self, year: int) -> list:
        """
        Get all pdf links from the page
        :return: list of pdf links
        """
        url = urljoin(self.url, str(year))
        log.info(f"Getting pdf links from {url}")
        page = requests.get(url)
        # regex findall pdf links similar to this https://j-w-a.or.jp/wp/wp-content/uploads/2021/02/2012全日本ジュニア.pdf
        pdf_links = re.findall(r"https://j-w-a.or.jp/wp/wp-content/uploads/\d{4}/\d{2}/\d{4}.*.pdf", page.text)
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

    def convert_pdf(self, pdf_path: str) -> None:
        """
        Convert pdf to csv
        :param pdf_path: path to pdf
        :return: None
        """
        # todo: finish this
        log.info(f"Converting {pdf_path} to csv")
        tabby = tabula.read_pdf(pdf_path, pages="all", multiple_tables=True)
        for tables in tabby:
            print(tables.to_csv())

    @staticmethod
    def __readable_pdf(pdf: PageObject) -> bool:
        """
        Checks if there is readable text in the pdf
        :return: boolean
        """
        if len(pdf.extract_text()) > 0:
            return True
        return False

    def sort_pdfs(self) -> None:
        """
        Sort pdfs into folders based on readable text
        :return: none
        """
        log.info(f"Sorting pdfs in {self.pdf_root}")
        folder_contents = listdir(self.pdf_root)
        folder_contents.remove("image")
        folder_contents.remove("readable_text")
        for pdf in folder_contents:
            log.info(f"Checking {pdf}")
            try:
                with open(path.join(self.pdf_root, pdf), "rb") as file:
                    pdf_reader = PdfReader(file)
                    if self.__readable_pdf(pdf_reader.pages[0]):
                        log.info(f"Moving {pdf} to readable folder")
                        move(path.join(self.pdf_root, pdf), path.join(self.pdf_root, "readable_text", pdf))
                    else:
                        log.info(f"Moving {pdf} to image folder")
                        move(path.join(self.pdf_root, pdf), path.join(self.pdf_root, "image", pdf))
            except Exception as e:
                log.error(f"Error with {pdf} - {e}")

if __name__ == '__main__':
    japan = Japan()
    # japan.sort_pdfs()
    japan.convert_pdf("/Users/euanmeston/PycharmProjects/OpenWeightlifting/python_tools/database_handler/japan_pdf/readable_text/2022　IC_W.pdf")