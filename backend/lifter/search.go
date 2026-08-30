package lifter

import (
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

// Rivals expects bigData to already be filtered (gender, federation, year) and
// deduped to one entry per lifter, ranked descending by total - the shape
// dbtools.FilterLifts's cache already produces. It just windows around nameStr.
func Rivals(nameStr string, bigData []*structs.Lift) (rivalResults structs.RivalsResult) {
	const WINDOW_SIZE = 3

	rivalResults.Total = len(bigData)

	targetIndex := -1
	for idx, lift := range bigData {
		if lift.Lifter.Name == nameStr {
			targetIndex = idx
			break
		}
	}

	if targetIndex == -1 {
		return
	}

	start := targetIndex - WINDOW_SIZE
	if start < 0 {
		start = 0
	}
	end := targetIndex + WINDOW_SIZE + 1
	if end > len(bigData) {
		end = len(bigData)
	}

	for i := start; i < end; i++ {
		rival := bigData[i]
		rivalResults.Rivals = append(rivalResults.Rivals, structs.Rival{
			Position:   i + 1,
			Total:      rival.Total,
			Lifter:     rival.Lifter.Name,
			Federation: rival.Event.Federation,
		})
	}

	return
}

// FetchLifts should use the exact string provided (case-sensitive) by NameSearch.
// If name.Disambiguation is set, only the matching lifter (by that stable index,
// see LifterRoster.AssignDisambiguation) is returned rather than every lifter
// who ever shared this name.
func FetchLifts(name structs.NameSearch, leaderboard *structs.LeaderboardData) (lifterData structs.LifterHistory) {
	lifterData.NameStr = name.NameStr
	for _, lift := range leaderboard.AllTotals {
		if lift.Lifter.Name != name.NameStr {
			continue
		}
		if name.Disambiguation != nil && lift.Lifter.Disambiguation != *name.Disambiguation {
			continue
		}
		lifterData.Lifts = append(lifterData.Lifts, lift)
	}
	return
}
