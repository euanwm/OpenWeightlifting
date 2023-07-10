package structs

import (
	"backend/enum"
	"backend/utilities"
	"log"
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
	datetime, _ := utilities.StringToDate(e.Date)
	eventYear, _, _ := datetime.Date()
	return eventYear == year
}

func (e Entry) WithinDates(startDate, endDate string) bool {
	if startDate == enum.ZeroDate && endDate == enum.MaxDate {
		return true
	}
	datetime, _ := utilities.StringToDate(e.Date)
	startDateTime, _ := utilities.StringToDate(startDate)
	endDateTime, _ := utilities.StringToDate(endDate)
	if datetime.After(startDateTime) && datetime.Before(endDateTime) {
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
		names = append(names, e.AllTotals[position].Name)
	}
	return
}

func (e AllData) ProcessNames() (names []string) {
	for _, lift := range e.Lifts {
		if !utilities.Contains(names, lift.Name) {
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
