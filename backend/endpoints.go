package main //nolint:typecheck

import (
	"backend/dbtools"
	"backend/enum"
	"backend/lifter"
	"backend/structs"
	"backend/utilities"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// LeaderboardData is a global variable that is used to hold the leaderboard data.
var LeaderboardData structs.LeaderboardData

// QueryCache is a global variable that is used to cache queries for the leaderboard endpoint.
var QueryCache dbtools.QueryCache

// kinda self-explanatory
var EventsData structs.EventsData

// LifterRoster is a global variable that is used to hold the lifter roster.
var LifterRoster structs.LifterRoster

// ServerTime godoc
//
//	@Summary	Checking the servers localtime
//	@Description	Returns the current server time.
//	@Tags			Utilities and Testing
//	@Produce		json
//	@Success		200	{object}	structs.ContainerTime
//	@Router			/time [get]
func ServerTime(c *gin.Context) {
	hour, mins, sec := time.Now().Clock()
	retStruct := structs.ContainerTime{Hour: hour, Min: mins, Sec: sec}
	c.JSON(http.StatusOK, retStruct)
}

// SearchName godoc
//
//		@Summary	Search through lifter names
//		@Schemes
//		@Description	Looks up a lifter by name and returns a list of possible matches. Requires a minimum of 3 characters.
//		@Tags			GET Requests
//	 @Param name query string true "Name to search for, minimum 3 characters"
//	 @Param limit query int false "Limit the number of results" default(50)
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.NameSearchResults
//		@Failure		204	{object}	nil
//		@Router			/search [get]
func SearchName(c *gin.Context) {
	maxResults, err := strconv.Atoi(c.Query("limit"))
	if err != nil || maxResults < 0 {
		log.Println("Failed to parse limit, defaulting to 50")
		maxResults = 50
	}
	if len(c.Query("name")) >= 3 {
		nameStr := c.Query("name")
		results := lifter.NameSearch(nameStr, LifterRoster)

		results.Total = len(results.Names)

		// todo: remove this and implement a proper solution
		if len(results.Names) > maxResults {
			results.Names = results.Names[:maxResults]
		}
		// todo: bug here as an empty struct is returned if no results are found higher up in the stack
		if len(results.Names) > 1 {
			c.JSON(http.StatusOK, results)
		} else if len(results.Names) == 1 {
			if len(results.Names[0].NameStr) > 0 {
				c.JSON(http.StatusOK, results)
			} else {
				c.JSON(http.StatusNoContent, nil)
			}
		}
	}
}

// LifterHistory godoc
//
//		@Summary	Retrieve a lifter's history
//		@Schemes
//		@Description	Pull a lifter's history by name. The name must be an exact match and can be checked using the search endpoint.
//		@Tags			GET Requests
//	 @Param name query string true "Name of the lifter, must be an exact match"
//	 @Param federation query string false "Federation to filter lifts by"
//	 @Param disamb query int false "Disambiguation index from the search endpoint, for names shared by multiple lifters"
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.LifterHistory
//	 @Failure		204	{object}	nil
//	 @Failure		400	{object}	nil
//		@Router			/history [get]
func LifterHistory(c *gin.Context) {
	name := c.Query("name")
	federation := c.Query("federation")
	lifterSearch := structs.NameSearch{NameStr: name, Federation: federation}

	if disambStr := c.Query("disamb"); disambStr != "" {
		disamb, err := strconv.Atoi(disambStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "disamb must be an integer"})
			return
		}
		lifterSearch.Disambiguation = &disamb
	}

	lifterDetails := lifter.FetchLifts(lifterSearch, &LeaderboardData)

	// todo: maybe refactor this to use a query struct, but I think a larger scale refactor is in order
	if len(federation) > 0 {
		lifterDetails.Lifts = dbtools.KeepFederationLifts(lifterDetails.Lifts, federation)
	}

	lifterDetails.Lifts = dbtools.SortDate(lifterDetails.Lifts)
	lifterDetails.Lifts = utilities.ReverseSlice(lifterDetails.Lifts)

	if len(lifterDetails.Lifts) != 0 {
		c.JSON(http.StatusOK, lifterDetails)
	} else if len(lifterDetails.Lifts) == 0 {
		c.JSON(http.StatusNoContent, nil)
	}
}

// Leaderboard godoc
//
//		@Summary	Main table on the index page
//		@Description	This is the used on the index page of the website and pulls the highest single lift for a lifter within the selected filter.
//		@Tags			GET Requests
//
//	 @Param start query int false "Position to begin from within the full query" default(0)
//	 @Param stop query int false "Position to stop at within the full query" default(50)
//	 @Param sortBy query string false "Sort by either total or sinclair" default(total)
//	 @Param federation query string false "Federation or country to filter by"
//	 @Param weightclass query string false "Weightclass to filter by" default(MALL)
//	 @Param year query int false "Year to filter by, mutually exclusive with startdate/enddate"
//	 @Param startdate query string false "Start date to filter from, mutually exclusive with year"
//	 @Param enddate query string false "End date to filter to, mutually exclusive with year"
//
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.LeaderboardResponse
//		@Failure		400	{object}	nil
//		@Router			/leaderboard [get]
func Leaderboard(c *gin.Context) {
	// There are 2 sorted leaderboards currently, sinclair and total. We default to total.
	sortby, exists := c.GetQuery("sortBy")
	if !exists {
		sortby = "total"
	}

	// If no federation is selected then we assume all federations
	federation, exists := c.GetQuery("federation")
	if !exists {
		federation = enum.ALLFEDS
	}

	// If no weight category is selected then we default to everyone
	weightclass, exists := c.GetQuery("weightclass")
	if !exists {
		weightclass = "MALL"
	}

	// Filter by year or within a certain range of dates
	year, yearExists := c.GetQuery("year") // todo: fix frontend filters and remove the 69 enum for all years
	if len(year) == 2 {
		year = ""
		yearExists = false
	}

	startDate, startDateExists := c.GetQuery("startdate")
	if !startDateExists {
		if !yearExists {
			startDate = enum.ZeroDate
		}
	}

	if startDateExists && yearExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Year and date ranges are exclusive"})
		return
	}

	endDate, endDateExists := c.GetQuery("enddate")
	if !endDateExists {
		if !yearExists {
			endDate = enum.MaxDate
		}
	}

	if endDateExists && yearExists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Year and date ranges are exclusive"})
		return
	}

	if yearExists && !startDateExists && !endDateExists {
		oneYear, err := strconv.Atoi(year)
		if err != nil {
			panic(err)
		}
		startDate = year + "-01-01"
		endDate = strconv.Itoa(oneYear+1) + "-01-01"
	}

	// Amount of results and positions to start at in the leaderboard
	start, exists := c.GetQuery("start")
	if !exists {
		start = "0"
	}
	startInt, err := strconv.Atoi(start)
	if err != nil {
		panic(err)
	}
	stop, exists := c.GetQuery("stop")
	if !exists {
		stop = "50"
	}
	stopInt, err := strconv.Atoi(stop)
	if err != nil {
		panic(err)
	}

	body := structs.LeaderboardPayload{
		Start:       startInt,
		Stop:        stopInt,
		SortBy:      sortby,
		Federation:  federation,
		WeightClass: weightclass,
		Year:        year,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	leaderboardData := LeaderboardData.Select(body.SortBy) // Selects either total or sinclair sorted leaderboard
	fedData := dbtools.FilterLifts(leaderboardData, body, dbtools.WeightClassList[body.WeightClass], &QueryCache)
	c.JSON(http.StatusOK, fedData)
}

// LeaderboardSearch godoc
//
//	@Summary	Find a lifter's position within the leaderboard
//	@Schemes
//	@Description	Checks whether a lifter appears within the results of a given leaderboard query and returns their position.
//	@Tags			POST Requests
//	@Param			request	body	structs.SearchLeaderboardRequest	true	"Lifter to search for and the leaderboard query to search within"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.SearchLeaderboardResult
//	@Failure		400	{object}	nil
//	@Router			/leaderboard/search [post]
func LeaderboardSearch(c *gin.Context) {
	var body structs.SearchLeaderboardRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := body.ActiveQuery.SetDefaults(c)
	if err != nil {
		log.Println("Error setting defaults for active query: ", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Broken query"})
		return
	}

	// Check that the lifter exists
	validLifterName := lifter.NameSearch(body.LifterData.NameStr, LifterRoster)

	if validLifterName.Total == 0 {
		c.JSON(http.StatusOK, gin.H{"error": "Name not in database"})
		return
	}

	// Filter by federation if required
	var filterByFed []structs.NameSearch
	if validLifterName.Total >= 1 && body.LifterData.Federation != "" {
		filterByFed = append(filterByFed, utilities.Filter(validLifterName.Names, func(m structs.NameSearch) bool {
			return m.Federation == body.LifterData.Federation
		})...)
	} else if validLifterName.Total > 1 {
		log.Println("Multiple lifter names found, but no federation specified")
		filterByFed = validLifterName.Names
	}

	finalLifter := filterByFed[0]

	leaderboardData := LeaderboardData.Select(body.ActiveQuery.SortBy)

	// Now we see if the name appears in the query
	leaderboardResult := structs.SearchLeaderboardResult{
		LifterData: body.LifterData,
		Position:   dbtools.LeaderboardPosition(leaderboardData, body.ActiveQuery, &QueryCache, finalLifter),
		Query:      body.ActiveQuery,
	}

	c.JSON(http.StatusOK, leaderboardResult)
}

// SimilarNameSearch godoc
//
//	@Summary	Fuzzy search for similar lifter names
//	@Schemes
//	@Description	Returns lifter names that are phonetically or typographically similar to the query. Handles IWF "LASTNAME Firstname" format automatically.
//	@Tags			GET Requests
//	@Param name query string true "Name to search for"
//	@Param federation query string true "Federation the lifter competes in"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.NameSimilarityResults
//	@Failure		204	{object}	nil
//	@Router			/search/similarity [get]
func SimilarNameSearch(c *gin.Context) {
	name := c.Query("name")
	federation := c.Query("federation")
	if len(name) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "name parameter is required"})
		return
	}
	nameSearch := structs.NameSearch{NameStr: name, Federation: federation}
	results := lifter.SimilarNames(nameSearch, &LifterRoster)
	if results.Total == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}
	c.JSON(http.StatusOK, results)
}

// Rival godoc
//
//	@Summary	Find a lifter's closest rivals by total
//	@Schemes
//	@Description	Returns the lifters ranked immediately around the given lifter, both within their federation (or country) and across all federations combined, for the current competition year.
//	@Tags			GET Requests
//	@Param			name	query	string	true	"Name of the lifter, must be an exact match"
//	@Param			sex		query	string	true	"Gender to filter by"
//	@Param			fed		query	string	false	"Federation or country to filter by, defaults to all federations"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.RivalsCombined
//	@Router			/rivals [get]
func Rival(c *gin.Context) {
	nameStr := c.Query("name")
	sexStr := c.Query("sex")
	fedStr := c.Query("fed")

	if len(fedStr) == 0 {
		fedStr = enum.ALLFEDS
	}

	leaderboardData := LeaderboardData.Select(enum.Total)

	var catStr string

	switch sexStr {
	case enum.Male:
		catStr = "MALL"
	case enum.Female:
		catStr = "FALL"
	default:
		log.Println("rivals:invalid sex:", sexStr)
		catStr = ""
	}

	query := structs.LeaderboardPayload{
		SortBy:      enum.Total,
		WeightClass: catStr,
		StartDate:   enum.CurrentYearFilter(),
		EndDate:     enum.NextYearFitler(),
		Stop:        len(leaderboardData),
	}

	query.Federation = fedStr
	fedFiltered := dbtools.FilterLifts(leaderboardData, query, dbtools.WeightClassList[catStr], &QueryCache)

	query.Federation = enum.ALLFEDS
	allFiltered := dbtools.FilterLifts(leaderboardData, query, dbtools.WeightClassList[catStr], &QueryCache)

	response := structs.RivalsCombined{
		FederationRivals: lifter.Rivals(nameStr, fedFiltered.Data),
		CombinedRivals:   lifter.Rivals(nameStr, allFiltered.Data),
	}

	c.JSON(http.StatusOK, response)
}

// Events godoc
//
//		@Summary	Fetch available event metadata within a set date range
//		@Schemes
//		@Description	Metadata shows the name, federation and date of the event along with the filename in the event_data folder.
//		@Tags			GET Requests
//	 @Param startdate query string false "Start of the date range to filter events by, defaults to 60 days ago"
//	 @Param enddate query string false "End of the date range to filter events by, defaults to today"
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.EventsList
//		@Router			/events/list [get]
func Events(c *gin.Context) {
	var response structs.EventsList
	const dateFormat = "2006-01-02"
	now := time.Now()
	startDate := c.DefaultQuery("startdate", now.AddDate(0, 0, -60).Format(dateFormat))
	endDate := c.DefaultQuery("enddate", now.Format(dateFormat))

	response.Events = EventsData.FetchEventWithinDate(startDate, endDate)
	c.JSON(http.StatusOK, response)
}

// SingleEvent godoc
//
//		@Summary	Fetch a single event
//		@Schemes
//		@Description	Fetch a single event, either by federation and event ID, or by federation and event name. When looking up by name, an optional date can be supplied to filter multi-day events loaded under a single event name.
//		@Tags			GET Requests
//	 @Param fed query string true "Federation of the event"
//	 @Param id query string false "ID of the event, required if name is not provided"
//	 @Param name query string false "Name of the event, required if id is not provided"
//	 @Param date query string false "Date to filter results by, only applicable when looking up by name"
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.EventResponse
//		@Failure		204	{object}	nil
//		@Router			/events [get]
func SingleEvent(c *gin.Context) {
	var response structs.EventResponse
	var federation, fedExists = c.GetQuery("fed")
	var csvID, idExists = c.GetQuery("id")
	var date, dateExists = c.GetQuery("date")
	var eventNameReq, nameExists = c.GetQuery("name")
	// federation and csvID are required
	if fedExists && idExists {
		response = EventsData.FetchEventByID(federation, csvID)
	} else if fedExists && nameExists {
		// federation and event name are required
		response = EventsData.FetchByEventName(eventNameReq)
		// date is optional, but I'd recommend it because I fucking said so and I can't be bothered explaining at 2323hrs on a Tuesday-cunting-night
		// only reason why it's even here is some federations load their multi-day events as such and not all on the same day
		if dateExists {
			response = response.FilterByDate(date)
		}
	}

	if len(response.Lifts) == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}
	c.JSON(http.StatusOK, response)
}
