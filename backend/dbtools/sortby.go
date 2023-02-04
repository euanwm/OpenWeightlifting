package dbtools

import (
	"backend/enum"
	"backend/lifter"
	"backend/structs"
	"backend/utilities"
	"sort"
	"time"
)

func removeFollowingLifts(bigData []structs.Entry) (filteredData []structs.Entry) {
	var names []string
	var position []int
	for i, d := range bigData {
		if utilities.Contains(names, d.Name) == false {
			position = append(position, i)
			names = append(names, d.Name)
		}
	}
	for _, posInt := range position {
		filteredData = append(filteredData, bigData[posInt])
	}
	return
}

//Filter - Returns a slice of structs relating to the selected filter selection
func Filter(bigData []structs.Entry, filterQuery structs.LeaderboardPayload, weightCat structs.WeightClass) (filteredData []structs.Entry) {
	for _, lift := range bigData {
		if filterQuery.Federation == enum.ALLFEDS {
			filterQuery.Federation = lift.Federation
		}
		if lift.Federation == filterQuery.Federation && lift.WithinWeightClass(filterQuery.Gender, weightCat) {
			linkedIG, igHandle := lifter.CheckUserList(lift.Name)
			if linkedIG {
				lift.Instagram = igHandle
			}
			filteredData = append(filteredData, lift)
		}
		if len(filteredData) >= filterQuery.Stop {
			filteredData = removeFollowingLifts(filteredData)
			if len(filteredData) >= filterQuery.Stop {
				return
			}
		}
	}
	filteredData = removeFollowingLifts(filteredData)
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

func SortLiftsBy(bigData []structs.Entry, sortBy string) (finalData []structs.Entry) {
	switch sortBy {
	case enum.Total:
		SortTotal(bigData)
	case enum.Sinclair:
		SortSinclair(bigData)
	}
	finalData = append(finalData, bigData...)
	return
}
