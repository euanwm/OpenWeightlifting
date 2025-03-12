from __future__ import annotations

import logging
from collections import defaultdict
from dataclasses import dataclass
from datetime import datetime
from itertools import chain
from pathlib import Path
from typing import Callable, Iterable, Iterator
from bisect import bisect_left

import requests
from bs4 import BeautifulSoup, Tag

from .result_dataclasses import Result
from .static_helpers import results_to_csv

logging.basicConfig()
logger = logging.getLogger(__name__)


def write_swiss_data(db_dir: str) -> None:
    """
    All swiss data is accessible from a single page. Visiting this page takes
    longer than transforming the data and writing some csv files. So until
    this is not the case anymore, we just replace all results with the scraped
    ones and don't have to concern ourselves with any update logic.
    """

    data = get_swiss_data()

    dir = Path(db_dir)

    dir.mkdir(exist_ok=True)

    for p in dir.glob("*"):
        if p.is_file():
            p.unlink()

    # sort the events by date to minimize diffs
    for i, results in enumerate(sorted(data.values(), key=lambda rs: rs[0].date)):
        if results:
            results_to_csv(db_dir, f"{i:03d}", results)


def get_html() -> str:
    url = "https://swiss-olympic-weightlifting.ch/resultats"

    try:
        response = requests.get(url)
        response.raise_for_status()
        return response.text
    except requests.exceptions.RequestException as e:
        logger.error(f"Error fetching HTML: {e}")
        return ""


def extract_entries(html: str) -> Iterator[Entry]:
    soup = BeautifulSoup(html, "html.parser")
    table = soup.find("tbody")
    if not isinstance(table, Tag):
        raise ValueError("Could not find result table in html")

    return filter_map(row_to_entry, table.find_all("tr"))


def get_swiss_data() -> dict[str, list[Result]]:
    html = get_html()
    entries = extract_entries(html)
    results = filter_map(entry_to_result, entries)
    events = group_results_by_event(results)

    return events


def group_results_by_event(results: Iterable[Result]) -> dict[str, list[Result]]:
    events = defaultdict(lambda: [])
    for e in results:
        events[e.event].append(e)
    return events


def entry_to_result(entry: Entry) -> Result | None:
    try:
        return Result(
            event=entry.event_name,
            date=entry.event_date.strftime("%Y-%m-%d"),
            category=cat_string(entry),
            lifter_name=f"{entry.lifter_firstname} {entry.lifter_lastname}",
            bodyweight=entry.bodyweight,
            snatch_1=entry.snatch_1,
            snatch_2=entry.snatch_2,
            snatch_3=entry.snatch_3,
            cj_1=entry.cj_1,
            cj_2=entry.cj_2,
            cj_3=entry.cj_3,
            best_snatch=entry.best_snatch,
            best_cj=entry.best_cj,
            total=entry.total,
        )
    except Exception as e:
        logger.info(f"Couldn't map {entry} because {e}", exc_info=True)
        return None


def mens_weightclass(weight: float) -> str:
    classes = [61, 67, 73, 81, 89, 96, 102, 109]
    labels = [f"{w}kg" for w in classes] + [f"{classes[-1]}+kg"]
    return labels[bisect_left(classes, weight)]


def womens_weightclass(weight: float) -> str:
    classes = [49, 55, 59, 64, 71, 76, 81, 87]
    labels = [f"{w}kg" for w in classes] + [f"{classes[-1]}+kg"]
    return labels[bisect_left(classes, weight)]


def is_male(entry: Entry) -> bool:
    if entry.category:
        return "m" in entry.category.lower()
    raise ValueError(f"Cannot determine gender of {entry}")


def cat_string(entry: Entry) -> str:
    if is_male(entry):
        return f"Men's Senior {mens_weightclass(entry.bodyweight)}"

    return f"Women's Senior {womens_weightclass(entry.bodyweight)}"


@dataclass
class Entry:
    """Corresponds to a table-row in https://swiss-olympic-weightlifting.ch/resultats"""

    event_date: datetime
    event_name: str
    lifter_firstname: str
    lifter_lastname: str
    category: str
    bodyweight: float
    club: str
    lifter_dob: datetime
    snatch_1: int
    snatch_2: int
    snatch_3: int
    best_snatch: int
    cj_1: int
    cj_2: int
    cj_3: int
    best_cj: int
    total: int
    sinclair: float
    place: str  # -.-

    def __post_init__(self):
        if not all(
            [
                self.event_name,
                self.lifter_firstname,
                self.lifter_firstname,
            ]
        ):
            raise ValueError("These fields can't be empty")


def row_to_entry(row: Tag) -> Entry | None:
    """
    The HTML table contains some annoying metadata-columns that
    we don't want to parse. We hardcode the indices of the interesting
    columns here and hope the table doesn't change
    """

    def parse_attempt(s: str) -> int:
        return int(s) if s else 0

    try:
        cells = list(map(lambda c: c.text.strip(), row.find_all("td")))
        return Entry(
            event_date=datetime.strptime(cells[0], "%d.%m.%Y"),
            event_name=cells[1],
            lifter_lastname=cells[8],
            lifter_firstname=cells[9],
            category=cells[10],
            bodyweight=float(cells[11]),
            club=cells[12],
            lifter_dob=datetime.strptime(cells[13], "%Y"),
            snatch_1=parse_attempt(cells[15]),
            snatch_2=parse_attempt(cells[16]),
            snatch_3=parse_attempt(cells[17]),
            best_snatch=parse_attempt(cells[18]),
            cj_1=parse_attempt(cells[19]),
            cj_2=parse_attempt(cells[20]),
            cj_3=parse_attempt(cells[21]),
            best_cj=parse_attempt(cells[22]),
            total=parse_attempt(cells[23]),
            sinclair=float(cells[24]),
            place=cells[25],
        )
    except Exception as e:
        # this is fine, some rows correspond to non-athlete results that
        # are not parseable as an Entry
        logger.debug(f"Couldn't parse {row} because {e}", exc_info=True)
        return None


# requires python 3.12
# def filter_map[T, U](f: Callable[[T], U | None], l: Iterable[T]) -> Iterator[U]:
def filter_map(f, l):
    return filter(None, map(f, l))


def main():
    print(f"Fetching data")
    data = get_swiss_data()
    print("Swiss events:")
    for name, results in data.items():
        print(f"  {name} with {len(results)} lifters")

    max_snatch = max(chain.from_iterable(data.values()), key=lambda e: e.best_snatch)
    print(
        f"Best snatch in Switzerland: {max_snatch.best_snatch} by {max_snatch.lifter_name}"
    )


if __name__ == "__main__":
    logger.setLevel(logging.INFO)
    main()
