package dbtools

import (
	"backend/enum"
	"backend/structs"
	"log"
	"sort"
)

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

func TopPerformance(bigData []structs.Entry, sortBy string, gender string) (finalData []structs.Entry) {
	log.Println("Building query " + sortBy + ":" + gender)
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
