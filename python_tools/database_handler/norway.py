""" API Tool for the Norwegian Weightlifting Federation """
import datetime
from requests import get
from urllib.parse import urljoin
from re import search

from .result_dataclasses import Result
from .static_helpers import load_json

import os

CatCodes = {
    "J": "Junior",
    "S": "Senior",
    "m": "Men's",
    "K": "Women's",
    "U": "Youth",
    "M": "Men's",
}

class Norway:
    def __init__(self):
        self.base_url: str = "https://nvf-backend.herokuapp.com/api/public/stevner/"
        self.results_root: str = "../backend/event_data/NVF"
        self.catlist = load_json(f"{os.getcwd()}/database_handler/gender_categories.json")


    def get_event_list(self) -> list[int]:
        """Returns a list of event IDs"""
        event_list: list[int] = []
        query = f"?fra-dato=2023-01-01&til-dato={self.__todays_date()}"
        # omg buh eror handlin
        res = get(urljoin(self.base_url, query)).json()
        # cope
        for event in res:
            event_list.append(event["id"])
        return event_list

    def fetch_event(self, event_id) -> list[Result]:
        res = get(urljoin(self.base_url, str(event_id))).json()
        results = res['puljer'][0]['resultater']
        comp_name = f"{res['klubbName']} {res['stevnetype']}"
        for x in results:
            try:
                cat_code = self.parse_cat_code(x['kategori']['forkortelse'], x['vektklasse']['navn'])
                datac = self.__assign_dataclass(x, cat_code, comp_name)
                datac_to_list = [x for x in datac.__dict__.values()]
                print(datac_to_list)
            except ValueError as e:
                print(e)


    def parse_cat_code(self, cat_code: str, weight: str) -> str|ValueError:
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

    @staticmethod
    def __assign_dataclass(result: dict, category: str, comp_name: str) -> Result:
        """Assigns the dataclass"""
        return Result(
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

    @staticmethod
    def __todays_date():
        return datetime.datetime.now().strftime("%Y-%m-%d")
