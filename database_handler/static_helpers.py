"""Exclusively static methods"""
from json import load


def load_json(filepath: str) -> dict:
    """Pass a filepath, get a dict"""
    with open(filepath, 'r', encoding='utf-8') as file:
        json_dict: dict = load(file)
    return json_dict
