""" As queries become more complex this will store and update them to speed up responses """
import os
import json
import csv
from time import perf_counter

from .result_dataclasses import Result, DatabaseEntry
from .static_helpers import load_result_csv_as_list, append_to_csv, load_csv_as_list


class QueryThis:
    """ Hail me for I am the query machine """
    query_folder = "queries/"
    results_root = "database_root/"

    def __init__(self):
        pass

    def generate_gender_cats(self):
        """Yep"""
        with open("database_handler/gender_categories.json",
                  "r", encoding="utf-8") as gender_cats:
            category_file = json.load(gender_cats)
        results_dir = os.listdir(self.results_root)
        for dirs in results_dir:
            fed_root = os.listdir(os.path.join(self.results_root, dirs))
            for file in fed_root:
                results = load_result_csv_as_list(os.path.join(self.results_root, dirs, file))
                new_cat_file = self.__pick_genders(results, category_file)
        with open("database_handler/gender_categories.json",
                  "w", encoding="utf-8") as new_gender_cats:
            new_gender_cats.write(json.dumps(new_cat_file, indent=4))

    def __pick_genders(self, results_list: list, cats: dict):
        """fuck"""
        for entry in results_list:
            entry_dc = Result(*entry)
            if "Men's" in entry_dc.category and entry_dc.category not in cats['male']:
                cats['male'].append(entry_dc.category)
            elif "Women's" in entry_dc.category and entry_dc.category not in cats['female']:
                cats['female'].append(entry_dc.category)
        return cats

    def __load_gender_cats(self) -> dict:
        """pish"""
        with open("database_handler/gender_categories.json", 'r', encoding='utf-8') as gender_cat_file:
            cat_dict: dict = json.load(gender_cat_file)
        return cat_dict

    def __assign_sex(self, category: str):
        """did you assume my biological sex? yes"""
        category_list = self.__load_gender_cats()
        if category in category_list['male'] or "Men's" in category:
            return "male"
        elif category in category_list['female'] or "Women's" in category:
            return "female"

    def collate_all_db(self):
        """Top 100 query, by total initially"""
        print("Collating big DB...")
        query_filename = "collated_db.csv"
        gender_cat = self.__load_gender_cats()
        with open(os.path.join(self.query_folder, query_filename), 'w', encoding='utf-8') as big_db:
            csv_writer = csv.writer(big_db)
            for country in os.listdir(self.results_root):  # UK / US / etc.
                for result in os.listdir((os.path.join(self.results_root, country))):
                    loaded_results = load_result_csv_as_list(os.path.join(self.results_root, country, result))
                    for single_result in loaded_results:
                        single_result.append(country)
                        if single_result[2] in gender_cat['male']:
                            single_result[2] = 'male'
                        elif single_result[2] in gender_cat['female']:
                            single_result[2] = 'female'
                        else:
                            print(f"Unknown category: {single_result}")
                        csv_writer.writerow(single_result)

    def separate_main_db(self):
        """splits by gender"""
        print("Separating DBs by gender...")
        main_db = os.path.join(self.query_folder, "collated_db.csv")
        with open(main_db, 'r', encoding='utf-8') as big_db:
            db_reader = csv.reader(big_db)
            for line in db_reader:
                if (entry := DatabaseEntry(*line)).gender == 'male':
                    append_to_csv(os.path.join(self.query_folder, "male.csv"), line)
                elif entry.gender == 'female':
                    append_to_csv(os.path.join(self.query_folder, "female.csv"), line)
                else:
                    print(entry)

    def sort_by_total(self, gender: str) -> None:
        """Sorts by total and also removes anyone with multiple entries"""
        print(f"Sorting {gender} totals...")
        timer_start = perf_counter()
        filpe: str = os.path.join(self.query_folder, gender + ".csv")
        db_list = load_csv_as_list(filpe)
        shite = sorted(db_list, key=lambda z: int(z[13]), reverse=True)
        # todo: this is a shite way to iterate and remove lifters with multiple (and smaller) entries
        for x in shite:
            for count, y in enumerate(shite):
                if x[3] == y[3] and x[14] == y[14] and x != y and x[13] >= y[13]:
                    shite.remove(y)
        for x_2 in shite:
            for count, y_2 in enumerate(shite):
                if x_2[3] == y_2[3] and x_2[14] == y_2[14] and x_2 != y_2 and x_2[13] >= y_2[13]:
                    shite.remove(y_2)
        for x in shite:
            append_to_csv(os.path.join(self.query_folder, f"top_total_{gender}.csv"), x)
        print(f"sort_by_total completed in {perf_counter() - timer_start} seconds")

    def sort_by_ratio(self, gender: str) -> None:
        """Bodyweight ratio to total"""
        print(f"Sorting {gender} ratios...")
        buffer: int = 110
        filpe: str = os.path.join(self.query_folder, gender + ".csv")
        db_list = load_csv_as_list(filpe)
        for index, line in enumerate(db_list):
            if float(line[4]) > 0:
                ratio = float(line[13]) / float(line[4])
                db_list[index].append(ratio)
            else:
                db_list[index].append(0)
        sort_it = self.__shit_sorter(db_list)
        for x in sort_it[:buffer:]:
            append_to_csv(os.path.join(self.query_folder, f"top_ratio_{gender}.csv"), x)

    @staticmethod
    def __shit_sorter(old_shite: list):
        """i hate it, you hate, we all hate it"""
        # todo: add in a method to choose the index number to sort by
        shite = sorted(old_shite, key=lambda z: z[15], reverse=True)
        for x in shite:
            for count, y in enumerate(shite):
                if x[3] == y[3] and x[14] == y[14] and x != y and x[15] >= y[15]:
                    shite.remove(y)
        for x_2 in shite:
            for count, y_2 in enumerate(shite):
                if x_2[3] == y_2[3] and x_2[14] == y_2[14] and x_2 != y_2 and x_2[15] >= y_2[15]:
                    shite.remove(y_2)
        return shite

    def create_lifter_index(self):
        """creates a csv file of all lifters"""
        query_filename = "lifter_names.csv"
        lifters = []
        with open(os.path.join(self.query_folder, query_filename), 'w', encoding='utf-8') as big_db:
            csv_writer = csv.writer(big_db)
            for country in os.listdir(self.results_root):  # UK / US / etc.
                for result in os.listdir((os.path.join(self.results_root, country))):
                    loaded_results = load_result_csv_as_list(os.path.join(self.results_root, country, result))
                    for single_result in loaded_results:
                        lifter_sex = self.__assign_sex(single_result[2])
                        # every name must be lower case to avoid caps mistakes
                        row = [single_result[3].lower(), lifter_sex, country]
                        if row not in lifters:
                            lifters.append(row)
            csv_writer.writerows(lifters)


if __name__ == '__main__':
    queerer = QueryThis()
    # queerer.compile_gender_cats()
    # queerer.generate_gender_cats()
    queerer.collate_all_db()
    queerer.separate_main_db()
    queerer.sort_by_total('female')
    queerer.sort_by_total('male')
    # queerer.sort_by_ratio("female")
    #queerer.create_lifter_index()