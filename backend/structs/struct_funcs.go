package structs

import (
	"backend/enum"
	"log"
)

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
