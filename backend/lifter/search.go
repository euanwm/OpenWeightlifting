package lifter

import (
	"backend/structs"
	"strings"
)

// NameSearch takes a partial string and returns a slice of positions within the AllNames slice that could be a match
func NameSearch(nameStr string, nameList *[]string) (namePositions []int) {
	nameStr = strings.ToLower(nameStr)
	for pos, name := range *nameList {
		if strings.Contains(strings.ToLower(name), nameStr) {
			namePositions = append(namePositions, pos)
		}
	}
	return
}

// FetchLifts should use the exact string provided (case-sensitive) by NameSearch
func FetchLifts(name structs.NameSearch, leaderboard *structs.LeaderboardData) (lifterData structs.LifterHistory) {
	lifterData.NameStr = name.NameStr
	// todo: make this nicer, it's a temporary fix to get rid of some dumber shit I did
	for _, lift := range leaderboard.MaleTotals {
		if lift.Name == name.NameStr {
			lifterData.Lifts = append(lifterData.Lifts, lift)
		}
	}
	for _, lift := range leaderboard.FemaleTotals {
		if lift.Name == name.NameStr {
			lifterData.Lifts = append(lifterData.Lifts, lift)
		}
	}
	return
}
