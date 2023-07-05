""" API Tool for the Norwegian Weightlifting Federation """
import datetime
import logging
import os

from typing import Any
from urllib.parse import urljoin
from re import search
from requests import get

from .result_dataclasses import Result
from .static_helpers import load_json, write_to_csv

CatCodes = {
    "J": "Junior",
    "S": "Senior",
    "m": "Men's",
    "K": "Women's",
    "U": "Youth",
    "M": "Men's",
}


class Norway:
    """ API Tool for the Norwegian Weightlifting Federation """
    def __init__(self):
        self.base_url: str = "https://nvf-backend.herokuapp.com/api/public/stevner/"
        self.results_root: str = "../backend/event_data/NVF"
        self.catlist = load_json(
            f"{os.getcwd()}/database_handler/gender_categories.json")

    def get_event_list(self) -> list[int]:
        """Returns a list of event IDs"""
        logging.info("Fetching event list")
        event_list: list[int] = []
        query = f"?fra-dato=2023-01-01&til-dato={self.__todays_date()}"
        # omg buh eror handlin
        res = get(urljoin(self.base_url, query), timeout=120).json()
        # cope
        for event in res:
            event_list.append(event["id"])
        return event_list

    def fetch_event(self, event_id) -> list[Any, Result]:
        """Returns a list of results for a given event ID"""
        logging.info("Fetching event %s", event_id)
        res = get(urljoin(self.base_url, str(event_id)), timeout=120).json()
        results = res['puljer'][0]['resultater']
        comp_name = f"{res['klubbName']} {res['stevnetype']}"
        event_results = [list(Result.__annotations__.keys())]
        for result in results:
            try:
                cat_code = self.parse_cat_code(
                    result['kategori']['forkortelse'], result['vektklasse']['navn'])
                datac = self.__assign_dataclass(result, cat_code, comp_name)
                datac_to_list = list(x for x in datac.__dict__.values())
                event_results.append(datac_to_list)
            except ValueError as ex:
                print(ex)
        return event_results

    def parse_cat_code(self, cat_code: str, weight: str) -> str | ValueError:
        """Returns the full category name"""
        if "+" in weight:
            weight = f"{weight[1:]}+"

        if search(r"\d{2}", cat_code):
            cat_1 = f"{CatCodes[cat_code[0]]} Masters"
            cat_2 = cat_code[1:]
            if "+" in cat_2:
                cat_2 = f"{cat_2[1:]}+"
        else:
            cat_1, cat_2 = CatCodes[cat_code[0]], CatCodes[cat_code[1]]
        cat_params: list[str] = [cat_1, cat_2, weight]
        all_params: list[str] = self.catlist['male'] + self.catlist['female']
        for param in all_params:
            if all(x in param for x in cat_params):
                return param

        print(f"Category not found: {cat_params} / {cat_code}")
        return ValueError(f"Category not found: {cat_params} / {cat_code}")

    def __assign_dataclass(self, result: dict, category: str,
                           comp_name: str) -> Result:
        """Assigns the dataclass"""
        datac = Result(
            event=comp_name,
            date=result['dato'],
            category=category,
            lifter_name=result['navn'],
            bodyweight=result["vektklasse"]['vektklasse'],
            snatch_1=result['rykk1'],
            snatch_2=result['rykk2'],
            snatch_3=result['rykk3'],
            cj_1=result['stot1'],
            cj_2=result['stot2'],
            cj_3=result['stot3'],
            best_snatch=result['besteRykk'],
            best_cj=result['besteStot'],
            total=result['total'],
        )

        for key, value in datac.__dict__.items():
            if value is None:
                if key in ['snatch_1', 'snatch_2',
                           'snatch_3', 'cj_1', 'cj_2', 'cj_3']:
                    setattr(datac, key, 0)
                if key == 'best_snatch':
                    setattr(datac, self.__best_snatch(datac))
                if key == 'best_cj':
                    setattr(datac, key, self.__best_cj(datac))
                if key == 'total':
                    setattr(datac, key, self.__calc_total(datac))

        return datac

    @staticmethod
    def __best_snatch(result: Result) -> float:
        """ Returns the best snatch """
        return max(result.snatch_1, result.snatch_2, result.snatch_3)

    @staticmethod
    def __best_cj(result: Result) -> float:
        """ Returns the best cj """
        return max(result.cj_1, result.cj_2, result.cj_3)

    @staticmethod
    def __calc_total(result: Result) -> float:
        """ Returns the total """
        return result.best_snatch + result.best_cj

    @staticmethod
    def __todays_date():
        return datetime.datetime.now().strftime("%Y-%m-%d")

    def update_results(self):
        """Updates the results"""
        logging.info("Updating results")
        event_list = self.get_event_list()
        result_db_ids = [int(x.split(".")[0])
                         for x in os.listdir(self.results_root)]
        for event_id in event_list:
            if event_id not in result_db_ids:
                event_results = self.fetch_event(event_id)
                write_to_csv(self.results_root, event_id, event_results)
