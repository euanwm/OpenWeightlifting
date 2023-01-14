package dbtools

import (
	"backend/enum"
	"backend/structs"
	"backend/utilities"
	"sort"
	"time"
)

//Filter - Returns a slice of structs relating to the selected filter selection
func Filter(bigData []structs.Entry, filterQuery structs.LeaderboardPayload, weightCat structs.WeightClass) (filteredData []structs.Entry) {
	var sliceLen = filterQuery.Stop - filterQuery.Start
	var indexCount int
	for _, lift := range bigData {
		if filterQuery.Federation == enum.ALLFEDS {
			filterQuery.Federation = lift.Federation
		}
		if lift.Federation == filterQuery.Federation && lift.WithinWeightClass(filterQuery.Gender, weightCat) && indexCount < filterQuery.Start {
			indexCount++
		} else if lift.Federation == filterQuery.Federation && lift.WithinWeightClass(filterQuery.Gender, weightCat) && indexCount >= filterQuery.Start {
			filteredData = append(filteredData, lift)
		} else if len(filteredData) == sliceLen {
			return
		}
	}
	if len(filteredData[filterQuery.Start:]) < sliceLen {
		return filteredData[filterQuery.Start:]
	}
	return filteredData[filterQuery.Start:filterQuery.Stop]
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

func TopPerformance(bigData []structs.Entry, sortBy string) (finalData []structs.Entry) {
	switch sortBy {
	case enum.Total:
		SortTotal(bigData)
	case enum.Sinclair:
		SortSinclair(bigData)
	}
	//bigData = dropBombs(bigData) may need this in the future
	var names []string
	var position []int
	for i, d := range bigData {
		if utilities.Contains(names, d.Name) == false {
			position = append(position, i)
			names = append(names, d.Name)
		}
	}
	for _, posInt := range position {
		finalData = append(finalData, bigData[posInt])
	}
	return
}

// dropBombs Removes people who failed to register a total
func dropBombs(bigData []structs.Entry) (newData []structs.Entry) {
	cutoffInt := 0
	for i, meh := range bigData {
		if meh.Total == 0 && cutoffInt == 0 {
			cutoffInt = i
		}
	}
	newData = bigData[:cutoffInt]
	return
}
