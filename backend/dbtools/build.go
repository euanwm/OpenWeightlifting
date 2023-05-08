package dbtools

import (
	"backend/enum"
	"backend/structs"
	"log"
)

func BuildDatabase(leaderboardTotal *structs.LeaderboardData) {
	log.Println("buildDatabase called...")
	bigData := CollateAll()
	allLifts, _ := ParseData(bigData)
	*leaderboardTotal = structs.LeaderboardData{
		AllTotals:    SortLiftsBy(allLifts.Lifts, enum.Total),
		AllSinclairs: SortLiftsBy(allLifts.Lifts, enum.Sinclair),
	}
	log.Println("Database READY")
}
