"""Dataclasses for how each federation formats it's event results"""
from dataclasses import dataclass
from typing import Any


@dataclass
class DatabaseEntry:
    """
    UNPACK IT BRO (*)\n
    This is the standard entry format for the main collated database.
    """
    event: Any
    date: Any
    gender: Any
    lifter_name: Any
    bodyweight: Any
    snatch_1: Any
    snatch_2: Any
    snatch_3: Any
    cj_1: Any
    cj_2: Any
    cj_3: Any
    best_snatch: Any
    best_cj: Any
    total: int
    country: Any

    def bw_ration(self) -> float:
        """Ratio bro"""
        return int(self.total) / float(self.bodyweight)

    def made_lifts(self) -> int:
        """Yeah...i'll add this in eventually"""

@dataclass
class UKUSResult:
    """
    UNPACK IT BRO (*)\n
    UK and US entities use the same event result format
    """
    event: Any
    date: Any
    category: Any
    lifter_name: Any
    bodyweight: float
    snatch_1: Any
    snatch_2: Any
    snatch_3: Any
    cj_1: Any
    cj_2: Any
    cj_3: Any
    best_snatch: Any
    best_cj: Any
    total: Any

    def bw_ration(self) -> float:
        """Ratio bro"""
        return int(self.total) / float(self.bodyweight)

    def made_lifts(self) -> int:
        """Yeah...i'll add this in eventually"""


if __name__ == '__main__':
    test_list = ["Vehement Elite Athletics Presents:  The 2021 Rising Tides Developmental Meet", "2021-04-03",
                 "Open Men's 89kg", "Jason Longfellow", 83.4, -115, 115, 0, 145, 0, 0, 115, 145, 260]
    test_dc = UKUSResult(*test_list)
    print(test_dc.bw_ration())
