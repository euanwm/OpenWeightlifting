package dbtools

import (
	"sort"
)

// SortTotal Descending order by entry total
func SortTotal(sliceStructs []Entry) {
	sort.Slice(sliceStructs, func(i, j int) bool {
		return sliceStructs[i].Total > sliceStructs[j].Total
	})
}

func OnlyTopBestTotal(bigData []Entry) (finalData []Entry) {
	bigData = dropBombs(bigData)
	var names []string
	var position []int
	for i, d := range bigData {
		if contains(names, d.Name) == false {
			position = append(position, i)
			names = append(names, d.Name)
		}
	}
	for _, position := range position {
		finalData = append(finalData, bigData[position])
	}
	return
}

func contains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

// dropBombs Removes people who failed to register a total
func dropBombs(bigData []Entry) (newData []Entry) {
	cutoffInt := 0
	for i, meh := range bigData {
		if meh.Total == 0 && cutoffInt == 0 {
			cutoffInt = i
		}
	}
	newData = bigData[:cutoffInt]
	return
}
