package lifter

import (
	"backend/enum"
	"backend/structs"
	"sort"
	"strings"
)

// NameSearch is similar to NameSearch but will also return names with their federation
func NameSearch(nameStr string, lifterRoster structs.LifterRoster) (nameResults structs.NameSearchResults) {
	nameStr = strings.ToLower(nameStr)
	nameResults = lifterRoster.Search(nameStr)

	if len(nameResults.Names) == 0 {
		nameResults.Names = append(nameResults.Names, structs.NameSearch{NameStr: "", Federation: ""})
	}

	return
}

// SimilarNames returns entries whose names are fuzzy-similar to nameDetails.NameStr.
// It uses token sort ratio (handles IWF "LASTNAME Firstname" vs "Firstname Lastname" format)
// combined with token-level Jaro-Winkler (typo tolerance and partial matches like Chris/Christopher)
// and Soundex phonetic boosting for varied spellings of the same sound.
// Results are sorted by descending score.
func SimilarNames(nameDetails structs.NameSearch, lifterRoster *structs.LifterRoster) (similarNames structs.NameSimilarityResults) {
	seen := make(map[string]bool)

	for _, entry := range lifterRoster.Lifters {
		// todo: gender is part of the Add() function for the LifterRoster struct
		// so we should probably run through that index map instead of this
		key := entry.Name + "|" + entry.PrimaryFederation
		if seen[key] {
			continue
		}
		seen[key] = true

		score := combinedNameScore(nameDetails.NameStr, entry.Name)
		if score >= similarityThreshold {
			similarNames.Names = append(similarNames.Names, structs.NameSimilarity{
				NameStr:    entry.Name,
				Federation: entry.PrimaryFederation,
				Score:      float32(score),
			})
			similarNames.Total++
		}
	}

	sort.SliceStable(similarNames.Names, func(i, j int) bool {
		return similarNames.Names[i].Score > similarNames.Names[j].Score
	})

	return
}

func Rivals(nameStr string, sex string, fed string, bigData []*structs.Lift) (rivalResults structs.RivalsResult) {
	const WINDOW_SIZE = 3

	var names []string
	var liftPtr *structs.Lift
	var liftPos []int
	var targetIndex = -1

	// First pass: collect first occurrence of each lifter (best performance since data is pre-sorted)
	seenNames := make(map[string]bool)

	for idx, lift := range bigData {
		liftPtr = bigData[idx]
		if liftPtr.Lifter.Gender == sex && lift.WithinYear(enum.CurrentYearInt()) && lift.Event.SelectedFederation(fed) {
			if !seenNames[lift.Lifter.Name] {
				seenNames[lift.Lifter.Name] = true
				names = append(names, lift.Lifter.Name)
				liftPos = append(liftPos, idx)
				if lift.Lifter.Name == nameStr {
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
			rivalResults.Rivals = append(rivalResults.Rivals, structs.Rival{
				Position:   i + 1,
				Total:      rival.Total,
				Lifter:     rival.Lifter.Name,
				Federation: rival.Event.Federation,
			})
		}
	}

	return
}

// FetchLifts should use the exact string provided (case-sensitive) by NameSearch
func FetchLifts(name structs.NameSearch, leaderboard *structs.LeaderboardData) (lifterData structs.LifterHistory) {
	lifterData.NameStr = name.NameStr
	for _, lift := range leaderboard.AllTotals {
		if lift.Lifter.Name == name.NameStr {
			lifterData.Lifts = append(lifterData.Lifts, lift)
		}
	}
	return
}
