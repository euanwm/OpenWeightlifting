""" Checks all CSV files within the backend/event_data folder and that it 
matches the Result dataclass """
import logging
import re
import sys
from os import getcwd, listdir
from os.path import join
from typing import Optional

from database_handler.result_dataclasses import Result
from database_handler.static_helpers import load_result_csv_as_list

event_data_path: str = "../event_data"


def check_db() -> None:
    """ To be used as part of a GitHub action to check the database files are up-to-date """
    print("Checking database files...")
    fed_dir = [fed for fed in listdir(event_data_path) if "." not in fed]
    if (arg_db := __single_database()) is not None:
        fed_dir = arg_db
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


def __single_database() -> Optional[list[str]]:
    """ Checks the args passed to the script from the makefile """
    if len(sys.argv) == 2 and len(sys.argv[1]) > 0:
        print(f"Checking {sys.argv[1]} database")
        db_path = [join(getcwd(), event_data_path, sys.argv[1])]
        return db_path
    else:
        return None


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
                        print(f"Date format incorrect for {entry}\nFile: {csv_filepath}")
                        pass_test = False
                    # check that the entry total is less than 500
                    if entry.total > 500:
                        print(f"Total format incorrect for {entry}\nFile: {csv_filepath}")
                        pass_test = False
                    # todo: maybe add this back in or make it optional
                    """"
                    # if a total is 0, then it's a DNF or DSQ
                    if entry.total != 0:
                        # check best snatch
                        if entry.best_snatch != max(0.0, max([entry.snatch_1, entry.snatch_2, entry.snatch_3])):
                            print(f"Best snatch incorrect for {entry}\nFile: {csv_filepath}\n")
                            pass_test = False
                        # check best clean & jerk
                        if entry.best_cj != max(0.0, max([entry.cj_1, entry.cj_2, entry.cj_3])):
                            print(f"Best clean-jerk incorrect for {entry}\nFile: {csv_filepath}\n")
                            pass_test = False
                    """
                    # check that the total not a negative value
                    if entry.total < 0:
                        print(f"Total incorrect for {entry}\nFile: {csv_filepath}\n")
                        pass_test = False
                    # check that the total is the sum of the best snatch and best clean & jerk
                    if entry.total > 0 and entry.total != entry.best_snatch + entry.best_cj:
                        print(f"Total incorrect for {entry}\nFile: {csv_filepath}\n")
                        pass_test = False
            except ValueError:
                pass_test = False
                print(f"Error in file: {csv_filepath}")
    return pass_test


def assign_dataclass(data: list) -> Result:
    """ Assigns the data to the Result dataclass """
    try:
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
