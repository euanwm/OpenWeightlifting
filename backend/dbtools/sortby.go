package dbtools

import (
	"backend/enum"
	"backend/structs"
	"sort"
)

//FilterFederation - Returns a slice of structs relating to the selected federation
func FilterFederation(bigData []structs.Entry, federation string, start int, stop int) (filteredData []structs.Entry) {
	var sliceLen = stop - start
	var indexCount int
	for _, lift := range bigData {
		if lift.Federation == federation && indexCount < start {
			indexCount++
		} else if lift.Federation == federation && indexCount >= start {
			filteredData = append(filteredData, lift)
		} else if len(filteredData) == sliceLen {
			return
		}
	}
	return nil
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

func TopPerformance(bigData []structs.Entry, sortBy string) (finalData []structs.Entry) {
	switch sortBy {
	case enum.Total:
		SortTotal(bigData)
	case enum.Sinclair:
		SortSinclair(bigData)
	}
	bigData = dropBombs(bigData)
	var names []string
	var position []int
	for i, d := range bigData {
		if Contains(names, d.Name) == false {
			position = append(position, i)
			names = append(names, d.Name)
		}
	}
	for _, position := range position {
		finalData = append(finalData, bigData[position])
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
