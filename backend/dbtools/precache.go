package dbtools

import (
	"backend/enum"
	"backend/structs"
)

// PreCacheQuery - Queries to pre-cache on startup. It's lazy, but combats the worst offending filter options by pre-caching them.
// I mean, you could literally pre-cache the whole fucking thing but as we all know from the previous sentence...I'm lazy.
func PreCacheQuery() (permutation []structs.LeaderboardPayload) {
	sortBy := []string{"total", "sinclair"}
	federation := []string{"allfeds", "UK", "US", "NVF", "AUS", "FFH", "IWF", "OPEN"}
	weightClass := []string{"MALL", "FALL"}
	// create all permutations and add them to the list
	for _, s := range sortBy {
		for _, f := range federation {
			for _, w := range weightClass {
				permutation = append(permutation, structs.LeaderboardPayload{
					SortBy:      s,
					Federation:  f,
					WeightClass: w,
					Year:        enum.AllYearsStr,
					StartDate:   enum.ZeroDate,
					EndDate:     enum.MaxDate,
					Start:       0,
					Stop:        50,
				})
			}
		}
	}
	return
}
