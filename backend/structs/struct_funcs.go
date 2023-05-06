package structs

import (
	"backend/enum"
	"backend/utilities"
	"log"
	"time"
)

func (e LifterHistory) GenerateChartData() ChartData {
	// todo: implement DRY principle
	var data ChartData
	for _, lift := range e.Lifts {
		data.Dates = append(data.Dates, lift.Date)
	}
	data.SubData = append(data.SubData, ChartSubData{
		Title:     "Competition Total",
		DataSlice: IterateFloatSlice(e.Lifts, enum.Total),
	})
	data.SubData = append(data.SubData, ChartSubData{
		Title:     "Best Snatch",
		DataSlice: IterateFloatSlice(e.Lifts, enum.BestSnatch),
	})
	data.SubData = append(data.SubData, ChartSubData{
		Title:     "Best C&J",
		DataSlice: IterateFloatSlice(e.Lifts, enum.BestCJ),
	})
	data.SubData = append(data.SubData, ChartSubData{
		Title:     "Bodyweight",
		DataSlice: IterateFloatSlice(e.Lifts, enum.Bodyweight),
	})
	return data
}

func (e Entry) WithinWeightClass(gender string, catData WeightClass) bool {
	if catData.Gender == enum.ALLCATS {
		return true
	}
	if catData.Gender == gender && catData.Upper >= e.Bodyweight && catData.Lower <= e.Bodyweight {
		return true
	}
	return false
}

func (e Entry) WithinYear(year int) bool {
	if year == enum.AllYears {
		return true
	}
	const rfc3339partial string = "T15:04:05Z"
	datetime, _ := time.Parse(time.RFC3339, e.Date+rfc3339partial)
	eventYear, _, _ := datetime.Date()
	if eventYear == year {
		return true
	}
	return false
}

func (e Entry) SelectedFederation(federation string) bool {
	if federation == enum.ALLFEDS {
		return true
	}
	if e.Federation == federation {
		return true
	}
	return false
}

func (e LeaderboardData) FetchNames(posSlice []int) (names []string) {
	for _, position := range posSlice {
		names = append(names, e.AllNames[position])
	}
	return
}

func (e AllData) ProcessNames() (names []string) {
	for _, lift := range e.Lifts {
		if utilities.Contains(names, lift.Name) == false {
			names = append(names, lift.Name)
		}
	}
	return
}

func (e LeaderboardData) Select(sortBy string) *[]Entry {
	switch sortBy {
	case enum.Total:
		return &e.AllTotals
	case enum.Sinclair:
		return &e.AllSinclairs
	}
	log.Println("LeaderboardData: Select - Error in selecting sinclair/total")
	return &[]Entry{}
}
