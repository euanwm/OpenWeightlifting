"""Fill this up with all the tools to generate/update databases and queries"""
import sys
from sys import argv
from datetime import datetime

from python_tools.database_handler import DBHandler
from python_tools.database_handler import AustraliaWeightlifting, InternationalWF
from python_tools.database_handler import QueryThis


class CLICommands:
    """boring shit, will probably realise we don't need this later"""
    def update(self, db_name):
        """updates all databases"""
        match db_name:
            case "iwf":
                iwf_db = InternationalWF("../backend/database_root/IWF")
                iwf_db.update_results()
            case "uk":
                uk_db = DBHandler("https://bwl.sport80.com/", "../backend/database_root/UK")
                uk_db.update_results(datetime.now().year)
            case "us":
                us_db = DBHandler("https://usaweightlifting.sport80.com/", "../backend/database_root/US")
                us_db.update_results(datetime.now().year)
            case "aus":
                aus_db = AustraliaWeightlifting()
                aus_db.update_db()
            case "all":
                year = datetime.now().year
                uk_db = DBHandler("https://bwl.sport80.com/", "../backend/database_root/UK")
                uk_db.update_results(year)
                us_db = DBHandler("https://usaweightlifting.sport80.com/", "../backend/database_root/US")
                us_db.update_results(year)
                aus_db = AustraliaWeightlifting()
                aus_db.update_db()
                iwf_db = InternationalWF("../backend/database_root/IWF")
                iwf_db.update_results()
            case _:
                sys.exit(f"database not found: {db_name}")

    def refresh_cache(self):
        """refreshes the query cache"""
        query_handler = QueryThis()
        query_handler.create_lifter_index()
        query_handler.create_gender_dbs()
        query_handler.sort_by_total('male')
        query_handler.sort_by_total('female')
        query_handler.create_comp_reference_index()


if __name__ == '__main__':
    commands = CLICommands()
    match argv[1]:
        case "--update":
            print(f"updating database: {argv[2]}")
            commands.update(argv[2])
        case "--refresh":
            print("refreshing...")
            commands.refresh_cache()
        case _:
            print("not a command")
