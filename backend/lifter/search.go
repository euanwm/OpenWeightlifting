package lifter

import (
	"backend/dbtools"
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
func NewNameSearch(nameStr string, nameList *[]structs.Entry) (nameResults structs.NameSearchResults) {
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

func Rivals(nameStr string, sex string, fed string, year int, bigData []structs.Entry) (rivalResults structs.RivalsResult) {
	const WINDOW_SIZE = 3

	var names []string
	var liftPtr *structs.Entry
	var liftPos []int
	var targetIndex int = -1

	// First pass: collect first occurrence of each lifter (best performance since data is pre-sorted)
	seenNames := make(map[string]bool)

	for idx, lift := range bigData {
		liftPtr = &bigData[idx]
		if dbtools.GetGender(liftPtr) == sex && lift.WithinYear(year) && lift.SelectedFederation(fed) {
			if !seenNames[lift.Name] {
				seenNames[lift.Name] = true
				names = append(names, lift.Name)
				liftPos = append(liftPos, idx)
				if lift.Name == nameStr {
					targetIndex = len(names) - 1
				}
				rivalResults.Total++
			}
		}
	}

	// If target found, create window of rivals
	if targetIndex != -1 {
		start := targetIndex - WINDOW_SIZE
		if start < 0 {
			start = 0
		}
		end := targetIndex + WINDOW_SIZE + 1
		if end > len(liftPos) {
			end = len(liftPos)
		}

		// Add rivals in the window
		for i := start; i < end; i++ {
			originalIdx := liftPos[i]
			rival := bigData[originalIdx]
			rivalResults.Rivals = append(rivalResults.Rivals, struct {
				Position   int
				Total      structs.WeightKg
				Gender     string
				Name       string
				Federation string
			}{
				Position:   i + 1,
				Total:      rival.Total,
				Gender:     rival.Gender,
				Name:       rival.Name,
				Federation: rival.Federation,
			})
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
