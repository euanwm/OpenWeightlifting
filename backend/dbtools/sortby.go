package dbtools

import (
	"backend/enum"
	"backend/structs"
	"backend/utilities"
	"log"
	"sort"
	"time"
)

func removeFollowingLifts(bigData []structs.Entry) (filteredData []structs.Entry) {
	var names []string
	var position []int
	for i, d := range bigData {
		if !utilities.Contains(names, d.Name) {
			position = append(position, i)
			names = append(names, d.Name)
		}
	}
	for _, posInt := range position {
		filteredData = append(filteredData, bigData[posInt])
	}
	return
}

// Filter - Returns a slice of structs relating to the selected filter selection
func Filter(bigData []structs.Entry, filterQuery structs.LeaderboardPayload, weightCat structs.WeightClass, cache *QueryCache) (filteredData structs.LeaderboardResponse) {
	exists, liftPositions := cache.CheckQuery(filterQuery)

	if exists {
		log.Println("Cache hit")
		filteredData.Data, filteredData.Size = fetchLifts(&bigData, liftPositions, filterQuery.Start, filterQuery.Stop)
		return
	}

	var liftPostions []int
	for idx, lift := range bigData {
		liftptr := &bigData[idx]
		if getGender(liftptr) == weightCat.Gender {
			if lift.SelectedFederation(filterQuery.Federation) && lift.WithinWeightClass(WeightClassList[filterQuery.WeightClass].Gender, weightCat) && lift.WithinDates(filterQuery.StartDate, filterQuery.EndDate) {
				liftPostions = append(liftPostions, idx)
			}
		}
	}

	cache.AddQuery(filterQuery, liftPostions)

	filteredData.Data, filteredData.Size = fetchLifts(&bigData, liftPostions, filterQuery.Start, filterQuery.Stop)
	return
}

// fetchLifts - Returns a slice of structs relating to the selected filter selection, it will also remove any duplicate entries.
func fetchLifts(bigData *[]structs.Entry, pos []int, start int, stop int) (lifts []structs.Entry, size int) {
	log.Println("Fetching...")
	for _, p := range pos {
		lifts = append(lifts, (*bigData)[p])
	}
	log.Println("Removing dupies")
	lifts = removeFollowingLifts(lifts)

	if stop > len(lifts) {
		stop = len(lifts)
	}

	if start > len(lifts) {
		start = len(lifts)
	}

	size = len(lifts)
	lifts = lifts[start:stop]
	log.Println("Fetched")
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
