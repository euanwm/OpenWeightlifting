package dbtools

import (
	"backend/enum"
	"backend/structs"
	"log"
)

func BuildDatabase() *structs.LeaderboardData {
	log.Println("buildDatabase called...")
	bigData := CollateAll()
	male, female, _ := SortGender(bigData) // Throwaway the unknown genders as they're likely really young kids
	leaderboardTotal := &structs.LeaderboardData{
		AllNames:        append(male.ProcessNames(), female.ProcessNames()...),
		MaleTotals:      SortLiftsBy(male.Lifts, enum.Total),
		FemaleTotals:    SortLiftsBy(female.Lifts, enum.Total),
		MaleSinclairs:   SortLiftsBy(male.Lifts, enum.Sinclair),
		FemaleSinclairs: SortLiftsBy(female.Lifts, enum.Sinclair),
	}
	log.Println("Database READY")
	return leaderboardTotal
}
