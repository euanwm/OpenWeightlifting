package structs

import (
	"backend/enum"
	"backend/utilities"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func (e Lift) WithinWeightClass(gender string, catData WeightClass) bool {
	if catData.Gender == enum.ALLCATS {
		return true
	}
	if catData.Gender == gender && catData.Upper.GreaterThanOrEqual(e.Bodyweight) && catData.Lower.LessThanOrEqual(e.Bodyweight) {
		return true
	}
	return false
}

func (e Lift) WithinYear(year int) bool {
	if year == enum.AllYears {
		return true
	}
	return e.Year() == year
}

func (e Lift) Year() int {
	datetime, _ := utilities.StringToDate(e.Event.Date)
	eventYear, _, _ := datetime.Date()
	return eventYear
}

func (e Lift) CurrentAge() int {
	if e.ageOnDay == 0 {
		return 0
	}
	currentYear := time.Now().Year()
	return currentYear - e.Year() + e.ageOnDay
}

func (e Lift) WithinDates(startDate, endDate string) bool {
	if startDate == enum.ZeroDate && endDate == enum.MaxDate {
		return true
	}
	datetime, _ := utilities.StringToDate(e.Event.Date)
	startDateTime, _ := utilities.StringToDate(startDate)
	endDateTime, _ := utilities.StringToDate(endDate)
	if datetime.After(startDateTime) && datetime.Before(endDateTime) {
		return true
	}
	return false
}

func (e *Lift) SetAgeOnDay(ageOnDay int) {
	e.ageOnDay = ageOnDay
}

func (e Event) SelectedFederation(federation string) bool {
	if federation == enum.ALLFEDS {
		return true
	}
	if e.Federation == federation {
		return true
	}
	return false
}

func (e LeaderboardData) Select(sortBy string) []*Lift {
	var lifts []*Lift
	switch sortBy {
	case enum.Total:
		lifts = e.AllTotals
	case enum.Sinclair:
		lifts = e.AllSinclairs
	default:
		log.Println("LeaderboardData: Select - Error in selecting sinclair/total")
		return []*Lift{}
	}

	return lifts
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

func (e LeaderboardResponse) FilterByDate(eventDate string) (newData []*Lift, newSize int) {
	for index, entry := range e.Data {
		if entry.Event.Date == eventDate {
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
	key := name + "|" + gender + "|" + federation + "|" + strconv.Itoa(lift.CurrentAge())
	if lifter, ok := e.index[key]; ok {
		lifter.Lifts = append(lifter.Lifts, lift)
		return lifter
	}
	lifter := &Lifter{
		Name:              name,
		Gender:            gender,
		CurrentAge:        lift.CurrentAge(),
		PrimaryFederation: federation,
		Lifts:             []*Lift{lift},
	}
	e.index[key] = lifter
	e.Lifters = append(e.Lifters, lifter)
	return lifter
}

func (e Lifter) IsMale() bool {
	return e.Gender == enum.Male
}

func (e LifterRoster) Search(nameStr string) (lifters NameSearchResults) {
	for _, lifter := range e.Lifters {
		if strings.Contains(strings.ToLower(lifter.Name), strings.ToLower(nameStr)) {
			lifters.Names = append(lifters.Names, NameSearch{NameStr: lifter.Name, Gender: lifter.Gender, CurrentAge: lifter.CurrentAge, Federation: lifter.PrimaryFederation})
			lifters.Total++
		}
	}
	return
}
