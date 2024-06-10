package dbtools

import (
	"backend/enum"
	"backend/structs"
	"backend/utilities"
	"log"
	"sort"
	"time"
)

// FilterLifts - Returns a slice of structs relating to the selected filter selection
func FilterLifts(bigData []structs.Entry, filterQuery structs.LeaderboardPayload, weightCat structs.WeightClass, cache *QueryCache) (filteredData structs.LeaderboardResponse) {
	exists, positions := cache.CheckQuery(filterQuery)

	if exists {
		filteredData.Data, filteredData.Size = fetchLifts(&bigData, positions, &filterQuery)
		return
	}

	if exists == false && positions != nil {
		log.Println("Partial query called")
		bigData = fetchLiftsAll(&bigData, positions)
	}

	var names []string
	var liftPtr *structs.Entry
	var liftPositions []int
	for idx, lift := range bigData {
		liftPtr = &bigData[idx]
		if getGender(liftPtr) == weightCat.Gender && !utilities.Contains(names, lift.Name) {
			if lift.SelectedFederation(filterQuery.Federation) && lift.WithinWeightClass(WeightClassList[filterQuery.WeightClass].Gender, weightCat) && lift.WithinDates(filterQuery.StartDate, filterQuery.EndDate) {
				liftPositions = append(liftPositions, idx)
				names = append(names, lift.Name)
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

func PreCacheFilter(bigData []structs.Entry, filterQuery structs.LeaderboardPayload, weightCat structs.WeightClass, cache *QueryCache) {
	var names []string
	var liftPtr *structs.Entry
	var liftPositions []int
	for idx, lift := range bigData {
		liftPtr = &bigData[idx]
		if getGender(liftPtr) == weightCat.Gender && !utilities.Contains(names, lift.Name) {
			if lift.SelectedFederation(filterQuery.Federation) && lift.WithinWeightClass(WeightClassList[filterQuery.WeightClass].Gender, weightCat) && lift.WithinDates(filterQuery.StartDate, filterQuery.EndDate) {
				liftPositions = append(liftPositions, idx)
				names = append(names, lift.Name)
			}
		}
	}
	cache.AddQuery(filterQuery, liftPositions)
}

// fetchLifts - Returns a slice of structs relating to the selected filter selection, it will also remove any duplicate entries.
func fetchLifts(bigData *[]structs.Entry, pos []int, query *structs.LeaderboardPayload) (lifts []structs.Entry, size int) {
	for _, p := range pos {
		lifts = append(lifts, (*bigData)[p])
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

func fetchLiftsAll(bigData *[]structs.Entry, pos []int) (lifts []structs.Entry) {
	for _, p := range pos {
		lifts = append(lifts, (*bigData)[p])
	}
	return
}

// SortSinclair Descending order by entry sinclair
func SortSinclair(sliceStructs []structs.Entry) {
	sort.Slice(sliceStructs, func(i, j int) bool {
		return sliceStructs[i].Sinclair > sliceStructs[j].Sinclair
	})
}

// SortTotal Descending order by entry total
func SortTotal(sliceStructs []structs.Entry) {
	sort.Slice(sliceStructs, func(i, j int) bool {
		return sliceStructs[i].Total > sliceStructs[j].Total
	})
}

// SortDate Ascending order by lift date
func SortDate(liftData []structs.Entry) []structs.Entry {
	const rfc3339partial string = "T15:04:05Z" // todo - manually subscribe to the RFC3339 string instead (?)
	sort.Slice(liftData, func(i, j int) bool {
		liftI, _ := time.Parse(time.RFC3339, liftData[i].Date+rfc3339partial)
		liftJ, _ := time.Parse(time.RFC3339, liftData[j].Date+rfc3339partial)
		return liftI.Before(liftJ)
	})
	return liftData
}

func SortLiftsBy(bigData []structs.Entry, sortBy string) (sortedData []structs.Entry) {
	switch sortBy {
	case enum.Total:
		SortTotal(bigData)
	case enum.Sinclair:
		SortSinclair(bigData)
	}
	sortedData = append(sortedData, bigData...)
	return
}

func KeepFederationLifts(bigData []structs.Entry, federation string) (filteredData []structs.Entry) {
	for _, lift := range bigData {
		if lift.Federation == federation {
			filteredData = append(filteredData, lift)
		}
	}
	return
}
