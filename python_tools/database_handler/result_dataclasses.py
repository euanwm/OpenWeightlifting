"""Dataclasses for how each federation formats its event results"""
from dataclasses import dataclass
from typing import Any


@dataclass
class IWFHeaders:
    """Standard headers for the IWF events index"""
    id: int
    name: str
    date: str
    location: str


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


@dataclass
class Result:
    """
    UNPACK IT BRO (*)\n
    UK and US entities use the same event result format
    """
    event: str
    date: str
    category: str
    lifter_name: str
    bodyweight: float
    snatch_1: float
    snatch_2: float
    snatch_3: float
    cj_1: float
    cj_2: float
    cj_3: float
    best_snatch: float = 0.0
    best_cj: float = 0.0
    total: float = 0.0

    def __post_init__(self):
        self.__catch_nones()
        self.best_snatch = self.__best_snatch()
        self.best_cj = self.__best_cj()
        self.total = self.__total()

    def __best_snatch(self):
        return max(0.0, self.snatch_1, self.snatch_2, self.snatch_3)

    def __best_cj(self):
        return max(0.0, self.cj_1, self.cj_2, self.cj_3)

    def __total(self):
        if self.best_snatch == 0.0 or self.best_cj == 0.0:
            return 0.0
        return self.best_snatch + self.best_cj

    def __catch_nones(self):
        for key, value in self.__dict__.items():
            if value is None:
                if key in ['snatch_1', 'snatch_2', 'snatch_3', 'cj_1', 'cj_2', 'cj_3']:
                    setattr(self, key, 0)
