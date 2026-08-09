package dbtools

import (
	"backend/enum"
	"backend/structs"
	"log"
)

func BuildDatabase(leaderboardTotal *structs.LeaderboardData, eventmetadata *structs.EventsData, lifterRoster *structs.LifterRoster) {
	log.Println("buildDatabase called...")
	allLifts := CollateAll(eventmetadata, lifterRoster)
	// allLifts, badLifts := ParseData(bigData)
	// log.Println("Unable to parse ", len(badLifts.Lifts), " lifts")
	*leaderboardTotal = structs.LeaderboardData{
		AllTotals:    SortLiftsBy(allLifts.Lifts, enum.Total),
		AllSinclairs: SortLiftsBy(allLifts.Lifts, enum.Sinclair),
	}
	log.Println("Database READY")
}
