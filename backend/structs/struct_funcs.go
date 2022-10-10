package structs

import (
	"backend/enum"
	"log"
)

func (e LeaderboardData) Query(sortBy string, gender string, start int, stop int) []Entry {
	switch sortBy {
	case enum.Total:
		switch gender {
		case enum.Male:
			return e.MaleTotals[start:stop]
		case enum.Female:
			return e.FemaleTotals[start:stop]
		default:
			log.Println("Some cunts being wild with it...")
		}
	case enum.Sinclair:
		switch sortBy {
		case enum.Male:
			return e.MaleSinclairs[start:stop]
		case enum.Female:
			return e.FemaleSinclairs[start:stop]
		default:
			log.Println("Some cunts being wild with it...")
		}
	}
	return nil
}
