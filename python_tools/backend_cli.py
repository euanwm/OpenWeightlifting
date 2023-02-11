"""Fill this up with all the tools to generate/update databases and queries"""
import sys
from sys import argv
from datetime import datetime

from python_tools.database_handler import DBHandler
from python_tools.database_handler import AustraliaWeightlifting, InternationalWF


class CLICommands:
    """boring shit, will probably realise we don't need this later"""
    def update(self, db_name):
        """updates all databases"""
        match db_name:
            case "iwf":
                iwf_db = InternationalWF("../backend/event_data/IWF")
                iwf_db.update_results()
            case "uk":
                uk_db = DBHandler("https://bwl.sport80.com/", "../backend/event_data/UK")
                uk_db.update_results(datetime.now().year)
            case "us":
                us_db = DBHandler("https://usaweightlifting.sport80.com/", "../backend/event_data/US")
                us_db.update_results(datetime.now().year)
            case "aus":
                aus_db = AustraliaWeightlifting()
                aus_db.update_db()
            case "all":
                year = datetime.now().year
                uk_db = DBHandler("https://bwl.sport80.com/", "../backend/event_data/UK")
                uk_db.update_results(year)
                us_db = DBHandler("https://usaweightlifting.sport80.com/", "../backend/event_data/US")
                us_db.update_results(year)
                aus_db = AustraliaWeightlifting()
                aus_db.update_db()
                iwf_db = InternationalWF("../backend/event_data/IWF")
                iwf_db.update_results()
            case _:
                sys.exit(f"database not found: {db_name}")


if __name__ == '__main__':
    commands = CLICommands()
    match argv[1]:
        case "--update":
            print(f"updating database: {argv[2]}")
            commands.update(argv[2])
        case _:
            print("not a command")
