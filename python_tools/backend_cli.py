"""Fill this up with all the tools to generate/update databases and queries"""
import logging
import sys
from sys import argv
from datetime import datetime

from database_handler import AustraliaWeightlifting, InternationalWF, Norway, FranceInterface, DBHandler


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
                iwf_db = InternationalWF("../event_data/IWF")
                iwf_db.update_results()
            case "uk":
                uk_db = DBHandler("https://bwl.sport80.com/",
                                  "../event_data/UK")
                uk_db.update_results(datetime.now().year)
            case "us":
                us_db = DBHandler(
                    "https://usaweightlifting.sport80.com/", "../event_data/US")
                us_db.update_results(datetime.now().year)
            case "aus":
                aus_db = AustraliaWeightlifting()
                aus_db.update_db()
                # leaving these here for debugging
                # aus_db.rebuild_db()
                # aus_db.add_single(15)
            case "ffh":
                france = FranceInterface()
                france.new_update_results()
            case "all":
                year = datetime.now().year
                print("Updating UK Database")
                uk_db = DBHandler("https://bwl.sport80.com/",
                                  "../event_data/UK")
                uk_db.update_results(year)
                print("Updating US Database")
                us_db = DBHandler(
                    "https://usaweightlifting.sport80.com/", "../event_data/US")
                us_db.update_results(year)
                print("Updating AWF Database")
                aus_db = AustraliaWeightlifting()
                aus_db.update_db()
                print("Updating IWF Database")
                iwf_db = InternationalWF("../event_data/IWF")
                iwf_db.update_results()
                print("Updating NVF Database")
                norway = Norway()
                norway.update_results()
                print("Updating France Database")
                france = FranceInterface()
                france.update_results()
            case _:
                sys.exit(f"database not found: {db_name}")

    def build(self, db_name):
        """builds a database"""
        match db_name:
            case "ffh":
                france = FranceInterface()
                france.new_build_database()
            case _:
                sys.exit(f"database not found: {db_name}")

if __name__ == '__main__':
    commands = CLICommands()
    match argv[1]:
        case "--update":
            print(f"updating database: {argv[2]}")
            commands.update(argv[2])
        case "--build":
            print(f"building database: {argv[2]}")
            commands.build(argv[2])
        case _:
            print("not a command")
