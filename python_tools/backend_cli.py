"""Fill this up with all the tools to generate/update databases and queries"""
import logging
import sys
from sys import argv
from datetime import datetime

from database_handler import DBHandler
from database_handler import AustraliaWeightlifting, InternationalWF, Norway


# pylint: disable=too-few-public-methods
# Only has one public method but this is fine
class CLICommands:
    """boring shit, will probably realise we don't need this later"""

    def update(self, db_name):
        """updates all databases"""
        match db_name:
            case "nvf":
                logging.info("Updating NVF Database")
                norway = Norway()
                norway.update_results()
            case "iwf":
                iwf_db = InternationalWF("../backend/event_data/IWF")
                iwf_db.update_results()
            case "uk":
                uk_db = DBHandler("https://bwl.sport80.com/",
                                  "../backend/event_data/UK")
                uk_db.update_results(datetime.now().year)
            case "us":
                us_db = DBHandler(
                    "https://usaweightlifting.sport80.com/", "../backend/event_data/US")
                us_db.update_results(datetime.now().year)
            case "aus":
                aus_db = AustraliaWeightlifting()
                aus_db.update_db()
            case "all":
                year = datetime.now().year
                print("Updating UK Database")
                uk_db = DBHandler("https://bwl.sport80.com/",
                                  "../backend/event_data/UK")
                uk_db.update_results(year)
                print("Updating US Database")
                us_db = DBHandler(
                    "https://usaweightlifting.sport80.com/", "../backend/event_data/US")
                us_db.update_results(year)
                print("Updating AWF Database")
                aus_db = AustraliaWeightlifting()
                aus_db.update_db()
                print("Updating IWF Database")
                iwf_db = InternationalWF("../backend/event_data/IWF")
                iwf_db.update_results()
                print("Updating NVF Database")
                norway = Norway()
                norway.update_results()
            case _:
                sys.exit(f"database not found: {db_name}")


if __name__ == '__main__':
    commands = CLICommands()
    if len(argv) > 1:
        match argv[1]:
            case "--update":
                if len(argv) > 2:
                    print(f"updating database: {argv[2]}")
                    commands.update(argv[2])
                else:
                    print("No database name provided")
            case _:
                print("not a command")
    else:
        print("No arguments provided")
