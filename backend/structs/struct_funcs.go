package structs

import (
	"backend/enum"
	"backend/utilities"
	"fmt"
	"log"
	"reflect"
	"strings"
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
			if entry.Sn1 > 0 {
				makemiss[0]++
			}
			if entry.Sn2 > 0 {
				makemiss[1]++
			}
			if entry.Sn3 > 0 {
				makemiss[2]++
			}
			if entry.Sn1 != 0 || entry.Sn2 != 0 || entry.Sn3 != 0 {
				numberOfLifts++
			}
		}
	case enum.CleanAndJerk:
		for _, entry := range e.Lifts {
			if entry.CJ1 > 0 {
				makemiss[0]++
			}
			if entry.CJ2 > 0 {
				makemiss[1]++
			}
			if entry.CJ3 > 0 {
				makemiss[2]++
			}
			if entry.CJ1 != 0 || entry.CJ2 != 0 || entry.CJ3 != 0 {
				numberOfLifts++
			}
		}
	}

	for _, lift := range makemiss {
		makeRates = append(makeRates, int(float32(lift)/float32(numberOfLifts)*100))
	}
	return
}

func (e LifterHistory) BestLift(lift string) float32 {
	var bestLift float32
	switch lift {
	case enum.Snatch:
		for _, entry := range e.Lifts {
			if entry.BestSn > bestLift {
				bestLift = entry.BestSn
			}
		}
	case enum.CleanAndJerk:
		for _, entry := range e.Lifts {
			if entry.BestCJ > bestLift {
				bestLift = entry.BestCJ
			}
		}
	case enum.Total:
		for _, entry := range e.Lifts {
			if entry.Total > bestLift {
				bestLift = entry.Total
			}
		}
	}
	return bestLift
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

func (e LeaderboardData) FetchByEventName(eventName string) (eventData []Entry) {
	for _, entry := range e.AllTotals {
		if entry.Event == eventName || strings.Contains(entry.Event, eventName) {
			eventData = append(eventData, entry)
		}
	}
	return
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

func (e LeaderboardResponse) FilterByDate(eventDate string) (newData []Entry, newSize int) {
	for index, entry := range e.Data {
		if entry.Date == eventDate {
			newData = append(newData, e.Data[index])
		}
	}
	return
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
