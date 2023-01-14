package structs

import (
	"backend/enum"
	"backend/utilities"
	"log"
)

func (e Entry) WithinWeightClass(gender string, catData WeightClass) bool {
	if catData.Gender == enum.ALLCATS {
		return true
	}
	if catData.Gender == gender && catData.Upper >= e.Bodyweight && catData.Lower <= e.Bodyweight {
		return true
	}
	return false
}

func (e LeaderboardData) FetchNames(posSlice []int) (names []string) {
	for _, position := range posSlice {
		names = append(names, e.AllNames[position])
	}
	return
}

func (e AllData) ProcessNames() (names []string) {
	for _, lift := range e.Lifts {
		if utilities.Contains(names, lift.Name) == false {
			names = append(names, lift.Name)
		}
	}
	return
}

func (e LeaderboardData) Query(sortBy string, gender string) []Entry {
	switch sortBy {
	case enum.Total:
		switch gender {
		case enum.Male:
			return e.MaleTotals
		case enum.Female:
			return e.FemaleTotals
		default:
			log.Println("Some cunts being wild with totals...")
		}
	case enum.Sinclair:
		switch gender {
		case enum.Male:
			return e.MaleSinclairs
		case enum.Female:
			return e.FemaleSinclairs
		default:
			log.Println("Some cunts being wild with sinclairs...")
		}
	}
	return nil
}
