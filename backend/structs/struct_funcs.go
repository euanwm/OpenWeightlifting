package structs

import (
	"backend/enum"
	"backend/utilities"
	"fmt"
	"log"
	"reflect"
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

func (e LifterHistory) GenerateStats() LifterStats {
	var stats LifterStats
	stats.BestSnatch = e.BestLift(enum.Snatch)
	stats.BestCJ = e.BestLift(enum.CleanAndJerk)
	stats.BestTotal = e.BestLift(enum.Total)
	stats.MakeRateSnatches = e.MakeRates(enum.Snatch)
	stats.MakeRateCJ = e.MakeRates(enum.CleanAndJerk)
	return stats
}

func (e LifterHistory) MakeRates(lift string) (makeRates []int) {
	makemiss := []int{0, 0, 0}
	numberOfLifts := 0
	switch lift {
	case enum.Snatch:
		for _, entry := range e.Lifts {
			if entry.Sn1.IsPositive() {
				makemiss[0]++
			}
			if entry.Sn2.IsPositive() {
				makemiss[1]++
			}
			if entry.Sn3.IsPositive() {
				makemiss[2]++
			}
			if !entry.Sn1.IsZero() || !entry.Sn2.IsZero() || !entry.Sn3.IsZero() {
				numberOfLifts++
			}
		}
	case enum.CleanAndJerk:
		for _, entry := range e.Lifts {
			if entry.CJ1.IsPositive() {
				makemiss[0]++
			}
			if entry.CJ2.IsPositive() {
				makemiss[1]++
			}
			if entry.CJ3.IsPositive() {
				makemiss[2]++
			}
			if !entry.CJ1.IsZero() || !entry.CJ2.IsZero() || !entry.CJ3.IsZero() {
				numberOfLifts++
			}
		}
	}

	for _, lift := range makemiss {
		makeRates = append(makeRates, int(float32(lift)/float32(numberOfLifts)*100))
	}
	return
}

func (e LifterHistory) BestLift(lift string) WeightKg {
	var bestLift WeightKg
	switch lift {
	case enum.Snatch:
		for _, entry := range e.Lifts {
			bestLift = bestLift.Max(entry.BestSn)
		}
	case enum.CleanAndJerk:
		for _, entry := range e.Lifts {
			bestLift = bestLift.Max(entry.BestCJ)
		}
	case enum.Total:
		for _, entry := range e.Lifts {
			bestLift = bestLift.Max(entry.Total)
		}
	}
	return bestLift
}

func (e Entry) WithinWeightClass(gender string, catData WeightClass) bool {
	if catData.Gender == enum.ALLCATS {
		return true
	}
	if catData.Gender == gender && catData.Upper.GreaterThanOrEqual(e.Bodyweight) && catData.Lower.LessThanOrEqual(e.Bodyweight) {
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

func (e Entry) DiscordPrint() (rawString string) {
	rawString += "```"
	keys := reflect.ValueOf(e)
	for i := 0; i < keys.NumField(); i++ {
		rawString += keys.Type().Field(i).Name + ": " + fmt.Sprintf("%v", keys.Field(i).Interface()) + "\n"
	}
	rawString += "```"
	return
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

func (e EventsMetaData) FetchEventFP(index int) (federation, filename string) {
	return e.Federation[index], e.ID[index]
}

func (e EventsMetaData) FetchEventByName(eventName string) (federation, filename string) {
	for index, name := range e.Name {
		if name == eventName {
			return e.Federation[index], e.ID[index]
		}
	}
	return "", ""
}

func (e EventsMetaData) FetchEventWithinDate(startDate, endDate string) (events []SingleEventMetaData) {
	startDateTime, _ := utilities.StringToDate(startDate)
	endDateTime, _ := utilities.StringToDate(endDate)
	for index, date := range e.Date {
		eventDateTime, _ := utilities.StringToDate(date)
		if eventDateTime.After(startDateTime) && eventDateTime.Before(endDateTime) {
			events = append(events, SingleEventMetaData{
				Name:       e.Name[index],
				Federation: e.Federation[index],
				Date:       e.Date[index],
				ID:         e.ID[index],
			})
		}
	}
	return
}
