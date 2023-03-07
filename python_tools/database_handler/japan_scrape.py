import csv

import requests
import re
import logging as log

import tabula

from urllib.parse import urljoin
from os import mkdir, listdir, path
from shutil import move
from dataclasses import dataclass

from PyPDF2 import PdfReader, PageObject
from pandas import DataFrame

log.basicConfig(level=log.INFO)

# Dataclass for PDFs
@dataclass
class PDF:
    """ Dataclass for PDF table details """
    name: str
    table_shapes: list


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
            _, columns = tables.shape
            print(columns)

    def cleanup_table(self, table: list) -> list[list]:
        """
        Clean up the table
        :param table: table to clean
        :return: cleaned table
        """
        log.info(f"Cleaning up table")
        for x in table[3:]:
            x = x.split(",")
            # remove empty strings or specific character strings
            x = [i for i in x if i != "" and i != "○"]
            print(x)
        return table

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

    def sort_pdf_by_table(self) -> None:
        """
        Sort pdfs into folders based on table columns
        :return: None
        """
        pdf_folder = path.join(self.pdf_root, "readable_text")
        pdfs = listdir(pdf_folder)
        for pdf in pdfs:
            log.info(f"Checking {pdf}")
            pdf_dataclass = PDF(name=pdf, table_shapes=[])
            with open(path.join(pdf_folder, pdf), "rb") as file:
                tables = tabula.read_pdf(file, pages="all", multiple_tables=True)
                for table in tables:
                    pdf_dataclass.table_shapes.append(list(table.shape))
                log.info(f"Found {len(pdf_dataclass.table_shapes)} tables")
                pdf_dataclass.table_shapes = self.__condense_table_shapes(pdf_dataclass.table_shapes)
                self.__append_table_results(pdf_dataclass)

    def __condense_table_shapes(self, table_shapes: list[list]) -> list[list]:
        """ Combine shapes that have the same column count by adding the row count together """
        log.info(f"Condensing table shapes")
        condensed_shapes = []
        for shape in table_shapes:
            if not self.__shape_column_exists(shape, condensed_shapes):
                condensed_shapes.append(shape)
            else:
                for s in condensed_shapes:
                    if shape[1] == s[1]:
                        s[0] += shape[0]
        # sort list by column count
        condensed_shapes.sort(key=lambda x: x[1])
        return condensed_shapes


    def __shape_column_exists(self, shape: list, shapes: list[list]) -> bool:
        """ Check if a shape with the same column count exists in a list of shapes """
        for s in shapes:
            if shape[1] == s[1]:
                return True
        return False


    def __append_table_results(self, pdf_tables: PDF) -> None:
        """ Append the PDF dataclass data to a line within a csv file """
        with open(path.join(self.pdf_root, "extraction", "table_results.csv"), "a", newline="") as file:
            writer = csv.writer(file)
            pdf_tables.table_shapes.insert(0, pdf_tables.name)
            writer.writerow(pdf_tables.table_shapes)

    def __libre_translate(self, text: str) -> str:
        """
        Translate text using libre translate
        :param text: text to translate
        :return: translated text
        """
        log.info(f"Translating")
        url = "https://libretranslate.com/translate"
        params = {
            "q": text,
            "source": "ja",
            "target": "en"
        }
        response = requests.post(url, json=params)
        if response.status_code != 200:
            log.error(f"Error translating text - {response.status_code}")
            return ""
        return response.json()["translatedText"]

    def test_func(self):
        big_tuple = [(9, 7), (12, 18), (9, 7), (15, 19), (9, 7), (18, 17), (9, 7), (17, 21), (9, 7), (13, 17), (9, 7), (14, 19), (9, 7), (12, 18), (8, 7), (7, 7), (15, 17), (9, 7), (11, 20), (9, 7), (15, 20), (9, 7), (8, 19), (11, 19), (9, 7), (9, 18), (8, 18), (5, 18)]
        big_tuple = [list(x) for x in big_tuple]
        print(big_tuple)
        print(self.__condense_table_shapes(big_tuple))

if __name__ == '__main__':
    japan = Japan()
    # japan.sort_pdfs()
    # japan.convert_pdf("/Users/euanmeston/PycharmProjects/OpenWeightlifting/python_tools/database_handler/japan_pdf/readable_text/2022　IC_W.pdf")
    japan.sort_pdf_by_table()
    # japan.test_func()