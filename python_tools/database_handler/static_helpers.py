"""Exclusively static methods"""
import csv

from json import load, dumps
from os.path import join

from .result_dataclasses import Result


def results_to_dict(big_list: list) -> list[dict]:
    """flat list of results into a dict format"""
    dict_list = []
    for x in big_list:
        entry = Result(*x)
        dict_list.append(entry.__dict__)
    return dict_list


def load_csv_as_list(filepath):
    """csv to list"""
    return_list: list = []
    with open(filepath, 'r', encoding='utf-8') as big_file:
        csv_reader = csv.reader(big_file)
        for line in csv_reader:
            return_list.append(line)
    return return_list


def write_to_csv(base_dir, filepath_name, data):
    """yes"""
    print(f"creating {filepath_name}.csv...")
    with open(join(base_dir, f"{filepath_name}.csv"), 'w', encoding='utf-8') as file_boi:
        csv_writer = csv.writer(file_boi)
        csv_writer.writerows(data)


def append_to_csv(filepath_name, data):
    """yes"""
    with open(filepath_name, 'a+', encoding='utf-8') as file_boi:
        csv_writer = csv.writer(file_boi)
        csv_writer.writerow(data)


def load_result_csv_as_list(filepath: str) -> list:
    """Stuff"""
    results_list: list = []
    with open(filepath, "r", encoding='utf-8') as results_file:
        csv_read = csv.reader(results_file)
        for lines in csv_read:
            results_list.append(lines)
    return results_list[1::]  # Drops the header line


def write_json_file(filename, json_data) -> None:
    """never gonna give you up"""
    with open(filename, 'w', encoding="utf_8") as file:
        file.write(dumps(json_data, indent=4))


def append_json_file(filename, json_data) -> None:
    """never gonna let you down"""
    with open(filename, 'a', encoding="utf_8") as file:
        file.write(dumps(json_data, indent=4))


def load_json(filepath: str) -> dict:
    """Pass a filepath, get a dict"""
    with open(filepath, 'r', encoding='utf-8') as file:
        json_dict: dict = load(file)
    return json_dict


def get_subdomain(url: str) -> str:
    """Gives you the subdomain for the sport80 page, should move this to the sport80 package"""
    # todo: convert this to a regex
    str_stripped = url.lstrip("https://").split('.')
    return str_stripped[0]
