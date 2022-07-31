"""Fill this up with all the tools to generate/update databases and queries"""

import argparse

import database_handler


def update():
    """updates all databases"""


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("--update", help="updates all root DB's (not query cache)",
                        action="update")
    args = parser.parse_args()
