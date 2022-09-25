""" As queries become more complex this will store and update them to speed up responses """
import os
import json

from .static_helpers import load_result_csv_as_list, load_csv_as_list, write_to_csv


class QueryThis:
    """ Hail me for I am the query machine """
    query_folder = "queries/"
    results_root = "database_root/"

    def __init__(self):
        pass

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

    def sort_by_total(self, gender: str) -> None:
        """Sorts by total and also removes anyone with multiple entries"""
        print(f"Sorting {gender} totals...")
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
        write_to_csv(self.query_folder, f"top_total_{gender}", shite)

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
        query_filename = "lifter_index"
        lifters = []
        for country in os.listdir(self.results_root):  # UK / US / etc.
            for result in os.listdir((os.path.join(self.results_root, country))):
                loaded_results = load_result_csv_as_list(os.path.join(self.results_root, country, result))
                for single_result in loaded_results:
                    lifter_sex = self.__assign_sex(single_result[2])
                    # every name must be lowercase to avoid caps mistakes
                    row = [single_result[3].lower(), lifter_sex, country]
                    if row not in lifters:
                        lifters.append(row)
        write_to_csv(self.query_folder, query_filename, lifters)

    def create_gender_dbs(self):
        """Runs through all the result files and separates into M/F files"""
        male_results = []
        female_results = []
        unknown_results = []
        for country in os.listdir(self.results_root):
            for result in os.listdir((os.path.join(self.results_root, country))):
                loaded_results = load_result_csv_as_list(os.path.join(self.results_root, country, result))
                for single_result in loaded_results:
                    lifter_sex = self.__assign_sex(single_result[2])
                    single_result.append(country)
                    match lifter_sex:
                        case 'male':
                            male_results.append(single_result)
                        case 'female':
                            female_results.append(single_result)
                        case _:
                            single_result.append(result)
                            unknown_results.append(single_result)
        write_to_csv(self.query_folder, "male", male_results)
        write_to_csv(self.query_folder, "female", female_results)
        write_to_csv(self.query_folder, "unknown", unknown_results)

    def create_comp_reference_index(self):
        """creates a reference index with comp name, country, and filename"""
        listy = []
        for country in os.listdir(self.results_root):
            for result in os.listdir((os.path.join(self.results_root, country))):
                loaded_results = load_result_csv_as_list(os.path.join(self.results_root, country, result))
                try:
                    append_this = [loaded_results[0][0], country, result]
                    listy.append(append_this)
                except IndexError:
                    print(loaded_results)
        write_to_csv(self.query_folder, "comp_index", listy)
