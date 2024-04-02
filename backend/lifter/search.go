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
	if len(names) == 0 {
		names = append(names, "")
	}
	return
}

// NewNameSearch is similar to NameSearch but will also return names with their federation
func NewNameSearch(nameStr string, nameList *[]structs.Entry) (nameResults structs.NameFedSearchResults) {
	nameStr = strings.ToLower(nameStr)
	for _, lift := range *nameList {
		if strings.Contains(strings.ToLower(lift.Name), nameStr) {
			nameResults.Names = append(nameResults.Names, []struct {
				Name       string
				Federation string
			}{{Name: lift.Name, Federation: lift.Federation}}...)
		}
	}
	if len(nameResults.Names) == 0 {
		nameResults.Names = append(nameResults.Names, struct {
			Name       string
			Federation string
		}{Name: "", Federation: ""})
	}

	// drop duplicates if the federation AND name match - it's messy but it works
	for i := 0; i < len(nameResults.Names); i++ {
		for j := i + 1; j < len(nameResults.Names); j++ {
			if nameResults.Names[i].Name == nameResults.Names[j].Name && nameResults.Names[i].Federation == nameResults.Names[j].Federation {
				nameResults.Names = append(nameResults.Names[:j], nameResults.Names[j+1:]...)
				j--
			}
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
