""" Checks all CSV files within the backend/event_data folder and that it 
matches the Result dataclass """
import logging
import re
import sys
from os import getcwd, listdir
from os.path import join

from database_handler.result_dataclasses import Result
from database_handler.static_helpers import load_result_csv_as_list

event_data_path: str = "../backend/event_data"


def check_db() -> None:
    """ To be used as part of a GitHub action to check the database files are up-to-date """
    print("Checking database files...")
    fed_dir = [fed for fed in listdir(event_data_path) if "." not in fed]
    pass_test = True
    for fed in fed_dir:
        print(f"Checking {join(getcwd(), fed)} database")
        filepath = join(getcwd(), event_data_path, fed)
        if not check_files(filepath):
            pass_test = False
    if not pass_test:
        print("TEST FAILED")
        sys.exit(1)  # Apparently needed for GitHub Actions
    elif pass_test:
        print("TEST PASSED")
        sys.exit(0)

def check_files(folder_path: str) -> bool:
    """check_files() checks all CSV files within a folder and that it matches the 
    Result dataclass"""""
    pass_test = True
    for file in listdir(folder_path):
        csv_filepath = join(folder_path, file)
        file_data = load_result_csv_as_list(csv_filepath)
        for entry in file_data:
            try:
                if entry := assign_dataclass(entry):
                    # check that entry date is YYYY-MM-DD
                    if not re.match(r"\d{4}-\d{2}-\d{2}", entry.date):
                        print(
                            f"Date format incorrect for {entry}\nFile: {csv_filepath}")
                        pass_test = False
                    # check that the entry total is less than 500
                    if entry.total > 500:
                        print(
                            f"Total format incorrect for {entry}\nFile: {csv_filepath}")
                        pass_test = False
            except ValueError:
                pass_test = False
                print(f"Error in file: {csv_filepath}")
    return pass_test


def assign_dataclass(data: list) -> Result:
    """ Assigns the data to the Result dataclass """
    try:
        def to_float(value):
                
        try:
            return float(value)
        except ValueError:
            # Return a default value (like 0.0) if the conversion fails
            return 0.0
            
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
        except ValueError as ex:
            print(f"Issue with {data}\nException raised: {ex}")
            raise ValueError from ex


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    check_db()
