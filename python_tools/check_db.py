""" Checks all CSV files within the backend/event_data folder and that it matches the Result dataclass """
import logging, re
from os import getcwd, listdir
from os.path import join

from database_handler.result_dataclasses import Result
from database_handler.static_helpers import load_result_csv_as_list

event_data_path: str = "../backend/event_data"


def check_db() -> None:
    """ To be used as part of a GitHub action to check the database files are up-to-date """
    logging.info("Checking database files...")
    fed_dir = [fed for fed in listdir(event_data_path) if "." not in fed]
    for fed in fed_dir:
        logging.info(f"Checking {join(getcwd(), fed)} database")
        filepath = join(getcwd(), event_data_path, fed)
        check_files(filepath)


def check_files(folder_path: str) -> None:
    for file in listdir(folder_path):
        file_data = load_result_csv_as_list(join(folder_path, file))
        for entry in file_data:
            entry = assign_dataclass(entry)
            # check that entry date is YYYY-MM-DD
            if not re.match(r"\d{4}-\d{2}-\d{2}", entry.date):
                logging.error(f"Date format incorrect for {entry}")
            # check that the entry total is less than 500
            if entry.total > 500:
                logging.error(f"Total format incorrect for {entry}")


def assign_dataclass(data: list) -> Result:
    """ Assigns the data to the Result dataclass """
    to_return = Result(
        event=data[0],
        date=str(data[1]),
        category=data[2],
        lifter_name=data[3],
        bodyweight=float(data[4]),
        snatch_1=float(data[5]),
        snatch_2=float(data[6]),
        snatch_3=float(data[7]),
        cj_1=float(data[8]),
        cj_2=float(data[9]),
        cj_3=float(data[10]),
        best_snatch=float(data[11]),
        best_cj=float(data[12]),
        total=float(data[13]),
    )
    return to_return

if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    check_db()
