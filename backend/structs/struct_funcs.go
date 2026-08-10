package structs

import (
	"backend/enum"
	"backend/utilities"
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
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

// ToEntry flattens a Lift (plus its linked Event/Lifter) into the legacy Entry
// shape, for consumers (leaderboard filtering, event-by-name lookup) not yet
// migrated to the Lift/Event/Lifter model.
func (l *Lift) ToEntry() Entry {
	return Entry{
		Event:      l.Event.Name,
		Date:       l.Event.Date,
		Gender:     l.Lifter.Gender,
		Name:       l.Lifter.Name,
		Bodyweight: l.Bodyweight,
		Sn1:        l.Sn1,
		Sn2:        l.Sn2,
		Sn3:        l.Sn3,
		CJ1:        l.CJ1,
		CJ2:        l.CJ2,
		CJ3:        l.CJ3,
		BestSn:     l.BestSn,
		BestCJ:     l.BestCJ,
		Total:      l.Total,
		Sinclair:   float32(l.Sinclair),
		Federation: l.Event.Federation,
	}
}

func (e LeaderboardData) FetchNames(posSlice []int) (names []string) {
	for _, position := range posSlice {
		names = append(names, e.AllTotals[position].Lifter.Name)
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
	var lifts []*Lift
	switch sortBy {
	case enum.Total:
		lifts = e.AllTotals
	case enum.Sinclair:
		lifts = e.AllSinclairs
	default:
		log.Println("LeaderboardData: Select - Error in selecting sinclair/total")
		return &[]Entry{}
	}
	entries := make([]Entry, len(lifts))
	for i, lift := range lifts {
		entries[i] = lift.ToEntry()
	}
	return &entries
}

func (e LeaderboardData) FetchByEventName(eventName string) (eventData []Entry) {
	for _, lift := range e.AllTotals {
		if lift.Event.Name == eventName || strings.Contains(lift.Event.Name, eventName) {
			eventData = append(eventData, lift.ToEntry())
		}
	}
	return
}

func (e EventsData) FetchEventWithinDate(startDate, endDate string) (events []Event) {
	startDateTime, _ := utilities.StringToDate(startDate)
	endDateTime, _ := utilities.StringToDate(endDate)
	for _, event := range e.Events {
		eventDateTime, _ := utilities.StringToDate(event.Date)
		if eventDateTime.After(startDateTime) && eventDateTime.Before(endDateTime) {
			events = append(events, *event)
		}
	}
	return
}

func (e EventsData) FetchEventByID(federation, csvID string) (response EventResponse) {
	for _, event := range e.Events {
		if event.Federation == federation && event.CSVID == csvID {
			response.Event = *event
			for _, lift := range event.Results {
				response.Lifts = append(response.Lifts, *lift)
			}
			return
		}
	}
	return
}

// FetchByEventName combines lifts from every event whose name matches (exact
// or substring), since some federations split one event across multiple
// per-day CSV files that share a name. Event metadata comes from the first match.
func (e EventsData) FetchByEventName(eventName string) (response EventResponse) {
	for _, event := range e.Events {
		if event.Name == eventName || strings.Contains(event.Name, eventName) {
			if response.Event.Name == "" {
				response.Event = *event
			}
			for _, lift := range event.Results {
				response.Lifts = append(response.Lifts, *lift)
			}
		}
	}
	return
}

// FilterByDate narrows an already-fetched EventResponse down to lifts from a
// single day, for the multi-day-event-under-one-name case.
func (e EventResponse) FilterByDate(singleDate string) (response EventResponse) {
	response.Event = e.Event
	for _, lift := range e.Lifts {
		if lift.Event.Date == singleDate {
			response.Lifts = append(response.Lifts, lift)
		}
	}
	return
}

func (c *LeaderboardPayload) SetDefaults(gin *gin.Context) (err error) {
	if c.SortBy == "" {
		c.SortBy = "total"
	}
	if c.Federation == "" {
		c.Federation = enum.ALLFEDS
	}
	if c.WeightClass == "" {
		c.WeightClass = "MALL"
	}
	var yearExists = len(c.Year) != 0
	if len(c.Year) == 2 {
		c.Year = ""
		yearExists = false
	}
	if c.StartDate == "" {
		if !yearExists {
			c.StartDate = enum.ZeroDate
		}
	}
	if c.StartDate != "" && yearExists {
		return fmt.Errorf("year and date ranges are exclusive")
	}
	if c.EndDate == "" && !yearExists {
		c.EndDate = enum.MaxDate
	}
	if yearExists && c.StartDate == "" && c.EndDate == "" {
		oneYear, err := strconv.Atoi(c.Year)
		if err != nil {
			return err
		}
		c.StartDate = c.Year + "-01-01"
		c.EndDate = strconv.Itoa(oneYear+1) + "-01-01"
	}
	c.Start = 0
	c.Stop = 0
	return nil
}

func (e LeaderboardResponse) FilterByDate(eventDate string) (newData []Entry, newSize int) {
	for index, entry := range e.Data {
		if entry.Date == eventDate {
			newData = append(newData, e.Data[index])
		}
	}
	return
}

// Add finds or creates the Lifter matching name+category+federation, attaches
// lift to it, and returns the Lifter pointer to store on lift.Lifter.
func (e *LifterRoster) Add(name, category, federation string, lift *Lift) *Lifter {
	if e.index == nil {
		e.index = make(map[string]*Lifter)
	}
	gender := enum.ClassifyGender(category)
	key := name + "|" + gender + "|" + federation
	if lifter, ok := e.index[key]; ok {
		lifter.Lifts = append(lifter.Lifts, lift)
		return lifter
	}
	lifter := &Lifter{
		Name:   name,
		Gender: gender,
		Lifts:  []*Lift{lift},
	}
	e.index[key] = lifter
	e.Lifters = append(e.Lifters, lifter)
	return lifter
}
