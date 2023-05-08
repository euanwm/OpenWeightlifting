package lifter

import (
	"backend/structs"
	"backend/utilities"
	"strings"
)

// NameSearch takes a partial string and returns a slice of positions within the AllNames slice that could be a match
func NameSearch(nameStr string, nameList *[]structs.Entry) (names []string) {
	nameStr = strings.ToLower(nameStr)
	for _, lift := range *nameList {
		if strings.Contains(strings.ToLower(lift.Name), nameStr) && !utilities.SliceContains(lift.Name, names) {
			names = append(names, lift.Name)
		}
	}
	return
}

// FetchLifts should use the exact string provided (case-sensitive) by NameSearch
func FetchLifts(name structs.NameSearch, leaderboard *structs.LeaderboardData) (lifterData structs.LifterHistory) {
	lifterData.NameStr = name.NameStr
	for _, lift := range leaderboard.AllTotals {
		if lift.Name == name.NameStr {
			lifterData.Lifts = append(lifterData.Lifts, lift)
		}
	}
	return
}
