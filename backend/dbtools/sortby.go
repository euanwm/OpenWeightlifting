package dbtools

import (
	"backend/enum"
	"backend/structs"
	"backend/utilities"
	"sort"
	"time"
)

// FilterLifts - Returns a slice of structs relating to the selected filter selection
func FilterLifts(bigData []*structs.Lift, filterQuery structs.LeaderboardPayload, weightCat structs.WeightClass, cache *QueryCache) (filteredData structs.LeaderboardResponse) {
	queryState, positions := cache.CheckQuery(filterQuery)

	switch queryState {
	case None:
		cache.InitQuery(filterQuery)
	case Working:
		state := cache.QueryStatus(filterQuery)
		for state == Working {
			time.Sleep(100 * time.Millisecond)
			state = cache.QueryStatus(filterQuery)
			if state == Completed {
				filteredData.Data, filteredData.Size = fetchLifts(bigData, positions, &filterQuery)
				return
			}
		}
	case Completed:
		filteredData.Data, filteredData.Size = fetchLifts(bigData, positions, &filterQuery)
		return
	default:
		// if you hit this, fuck you
		panic("Invalid query state")
	}

	var names []string
	var liftPtr *structs.Lift
	var liftPositions []int
	for idx, lift := range bigData {
		liftPtr = bigData[idx]
		if liftPtr.Lifter.Gender == weightCat.Gender && !utilities.Contains(names, lift.Lifter.Name) {
			if lift.Event.SelectedFederation(filterQuery.Federation) && lift.WithinWeightClass(WeightClassList[filterQuery.WeightClass].Gender, weightCat) && lift.WithinDates(filterQuery.StartDate, filterQuery.EndDate) {
				liftPositions = append(liftPositions, idx)
				names = append(names, lift.Lifter.Name)
				filteredData.Data = append(filteredData.Data, lift)
			}
		}
	}
	cache.AddQuery(filterQuery, liftPositions)

	if filterQuery.Stop > len(liftPositions) {
		filterQuery.Stop = len(liftPositions)
	}

	if filterQuery.Start > len(liftPositions) {
		filterQuery.Start = len(liftPositions)
	}

	filteredData.Size = len(liftPositions)
	filteredData.Data = filteredData.Data[filterQuery.Start:filterQuery.Stop]
	return
}

func PreCacheFilter(bigData []*structs.Lift, filterQuery structs.LeaderboardPayload, weightCat structs.WeightClass, cache *QueryCache) {
	queryState, _ := cache.CheckQuery(filterQuery)

	switch queryState {
	case None:
		cache.InitQuery(filterQuery)
	case Working:
		state := cache.QueryStatus(filterQuery)
		for state == Working {
			time.Sleep(100 * time.Millisecond)
			state = cache.QueryStatus(filterQuery)
			if state == Completed {
				return
			}
		}
	case Completed:
		return
	default:
		// if you hit this, fuck you
		panic("Invalid query state")
	}

	var names []string
	var liftPtr *structs.Lift
	var liftPositions []int
	for idx, lift := range bigData {
		liftPtr = bigData[idx]
		if liftPtr.Lifter.Gender == weightCat.Gender && !utilities.Contains(names, lift.Lifter.Name) {
			if lift.Event.SelectedFederation(filterQuery.Federation) && lift.WithinWeightClass(WeightClassList[filterQuery.WeightClass].Gender, weightCat) && lift.WithinDates(filterQuery.StartDate, filterQuery.EndDate) {
				liftPositions = append(liftPositions, idx)
				names = append(names, lift.Lifter.Name)
			}
		}
	}
	cache.AddQuery(filterQuery, liftPositions)
}

func LeaderboardPosition(bigData []*structs.Lift, filterQuery structs.LeaderboardPayload, cache *QueryCache, lifterQuery structs.NameSearch) int {
	queryState, positions := cache.CheckQuery(filterQuery)

	switch queryState {
	case None:
		// cache.InitQuery(filterQuery)
		panic("Fuck you")
	case Working:
		state := cache.QueryStatus(filterQuery)
		for state == Working {
			time.Sleep(100 * time.Millisecond)
			state = cache.QueryStatus(filterQuery)
			if state == Completed {
				return lifterPosition(bigData, positions, lifterQuery)
			}
		}
	case Completed:
		return lifterPosition(bigData, positions, lifterQuery)
	default:
		// if you hit this, fuck you
		panic("Invalid query state")
	}
	return 0
}

func lifterPosition(bigData []*structs.Lift, pos []int, lifter structs.NameSearch) int {
	for i, d := range pos {
		liftData := bigData[d]
		if liftData.Lifter.Name == lifter.NameStr && liftData.Event.Federation == lifter.Federation {
			return i + 1
		}
	}
	return 0
}

// fetchLifts - Returns a slice of structs relating to the selected filter selection, it will also remove any duplicate entries.
func fetchLifts(bigData []*structs.Lift, pos []int, query *structs.LeaderboardPayload) (lifts []*structs.Lift, size int) {
	for _, p := range pos {
		lifts = append(lifts, bigData[p])
	}

	if query.Stop > len(lifts) {
		query.Stop = len(lifts)
	}

	if query.Start > len(lifts) {
		query.Start = len(lifts)
	}

	size = len(lifts)
	lifts = lifts[query.Start:query.Stop]
	return
}

// SortSinclair Descending order by lift sinclair
func SortSinclair(sliceStructs []*structs.Lift) {
	sort.Slice(sliceStructs, func(i, j int) bool {
		return sliceStructs[i].Sinclair.GreaterThan(sliceStructs[j].Sinclair)
	})
}

// SortTotal Descending order by lift total
func SortTotal(sliceStructs []*structs.Lift) {
	sort.Slice(sliceStructs, func(i, j int) bool {
		return sliceStructs[i].Total.GreaterThan(sliceStructs[j].Total)
	})
}

// SortDate Ascending order by entry date
func SortDate(liftData []*structs.Lift) []*structs.Lift {
	const rfc3339partial string = "T15:04:05Z" // todo - manually subscribe to the RFC3339 string instead (?)
	sort.Slice(liftData, func(i, j int) bool {
		liftI, _ := time.Parse(time.RFC3339, liftData[i].Event.Date+rfc3339partial)
		liftJ, _ := time.Parse(time.RFC3339, liftData[j].Event.Date+rfc3339partial)
		return liftI.Before(liftJ)
	})
	return liftData
}

// SortLiftsBy sorts a copy of bigData so callers using the same underlying
// slice for multiple sort orders (e.g. Total then Sinclair) don't clobber
// each other's already-returned results.
func SortLiftsBy(bigData []*structs.Lift, sortBy string) (sortedData []*structs.Lift) {
	sortedData = append(sortedData, bigData...)
	switch sortBy {
	case enum.Total:
		SortTotal(sortedData)
	case enum.Sinclair:
		SortSinclair(sortedData)
	}
	return
}

func KeepFederationLifts(bigData []*structs.Lift, federation string) (filteredData []*structs.Lift) {
	for _, lift := range bigData {
		if lift.Event.Federation == federation {
			filteredData = append(filteredData, lift)
		}
	}
	return
}
