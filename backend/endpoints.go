package main //nolint:typecheck

import (
	"backend/dbtools"
	"backend/discordbot"
	"backend/enum"
	"backend/lifter"
	"backend/structs"
	"backend/utilities"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

// DiscoKaren is a global variable that is used to hold the discord bot session.
var DiscoKaren discordbot.DiscordBot

// LeaderboardData is a global variable that is used to hold the leaderboard data.
var LeaderboardData structs.LeaderboardData

// this is remnant of the instagram linking code
// var lifterData = lifter.Build()

// QueryCache is a global variable that is used to cache queries for the leaderboard endpoint.
var QueryCache dbtools.QueryCache

// EventsData is a global variable that is used to hold the event metadata.
var EventsData structs.EventsMetaData

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
//	 @Param name query string true "Name to search for"
//	 @Param limit query int false "Limit the number of results"
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.NameSearchResults
//		@Router			/search [get]
func SearchName(c *gin.Context) {
	maxResults, err := strconv.Atoi(c.Query("limit"))
	if err != nil || maxResults < 0 {
		log.Println("Failed to parse limit, defaulting to 50")
		maxResults = 50
	}
	if len(c.Query("name")) >= 3 {
		nameStr := c.Query("name")
		results := lifter.NewNameSearch(nameStr, &LeaderboardData.AllTotals)

		results.Total = len(results.Names)

		// todo: remove this and implement a proper solution
		if len(results.Names) > maxResults {
			results.Names = results.Names[:maxResults]
		}
		c.JSON(http.StatusOK, results)
	}
}

// LifterGraph godoc
//
//		@Summary	Retrieve a lifter's record for use with ChartJS on the leaderboard page
//		@Schemes
//		@Description	This is used within the lifter page to display a lifter's record. It returns a JSON object that can be used with ChartJS without having to do any additional processing.
//		@Tags			GET Requests
//	 @Param name body string true "name"
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.ChartData
//	 @Failure		204	{object}	nil
//		@Router			/lifter [post]
func LifterGraph(c *gin.Context) {
	name := c.Query("name")
	federation := c.Query("federation")
	lifterSearch := structs.NameSearch{NameStr: name, Federation: federation}

	lifterDetails := lifter.FetchLifts(lifterSearch, &LeaderboardData)

	// todo: maybe refactor this to use a query struct, but I think a larger scale refactor is in order
	if len(federation) > 0 {
		lifterDetails.Lifts = dbtools.KeepFederationLifts(lifterDetails.Lifts, federation)
	}

	lifterDetails.Lifts = dbtools.SortDate(lifterDetails.Lifts)
	finalPayload := lifterDetails.GenerateChartData()
	if len(lifterDetails.Lifts) != 0 {
		c.JSON(http.StatusOK, finalPayload)
	} else if len(lifterDetails.Lifts) == 0 {
		c.JSON(http.StatusNoContent, nil)
	}
}

// LifterHistory godoc
//
//		@Summary	Retrieve a lifter's history
//		@Schemes
//		@Description	Pull a lifter's history by name. The name must be an exact match and can be checked using the search endpoint.
//		@Tags			GET Requests
//	 @Param name body string true "name"
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.LifterHistory
//	 @Failure		204	{object}	nil
//		@Router			/history [get]
func LifterHistory(c *gin.Context) {
	name := c.Query("name")
	federation := c.Query("federation")
	lifterSearch := structs.NameSearch{NameStr: name, Federation: federation}

	lifterDetails := lifter.FetchLifts(lifterSearch, &LeaderboardData)

	// todo: maybe refactor this to use a query struct, but I think a larger scale refactor is in order
	if len(federation) > 0 {
		lifterDetails.Lifts = dbtools.KeepFederationLifts(lifterDetails.Lifts, federation)
	}

	lifterDetails.Lifts = dbtools.SortDate(lifterDetails.Lifts)
	lifterDetails.Graph = lifterDetails.GenerateChartData()
	lifterDetails.Lifts = utilities.ReverseSlice(lifterDetails.Lifts)
	lifterDetails.Stats = lifterDetails.GenerateStats()

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
//	 @Param start query int false "Position to begin from within the full query"
//	 @Param stop query int false "Position to stop at within the full query"
//	 @Param sortby query string false "Sort by either total or sinclair"
//	 @Param federation query string false "Federation or country to filter by"
//	 @Param weightclass query string false "Weightclass to filter by"
//	 @Param year query int false "Year to filter by"
//	 @Param startdate query string false "Not currently used"
//	 @Param enddate query string false "Not currently used"
//
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.LeaderboardResponse
//		@Router			/leaderboard [post]
func Leaderboard(c *gin.Context) {
	sortby, exists := c.GetQuery("sortby")
	if !exists {
		sortby = "total"
	}
	federation, exists := c.GetQuery("federation")
	if !exists {
		federation = enum.ALLFEDS
	}
	weightclass, exists := c.GetQuery("weightclass")
	if !exists {
		weightclass = "MALL"
	}
	year, exists := c.GetQuery("year")
	if !exists {
		year = strconv.Itoa(enum.AllYears)
	}

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
	}

	// todo: remove this once the frontend filters have been updated to suit
	switch body.Year {
	case strconv.Itoa(enum.AllYears):
		body.StartDate = enum.ZeroDate
		body.EndDate = enum.MaxDate
	default:
		body.StartDate = body.Year + "-01-01"
		oneYear, err := strconv.Atoi(body.Year)
		if err != nil {
			panic(err)
		}
		body.EndDate = strconv.Itoa(oneYear+1) + "-01-01"
	}

	leaderboardData := LeaderboardData.Select(body.SortBy) // Selects either total or sinclair sorted leaderboard
	fedData := dbtools.FilterLifts(*leaderboardData, body, dbtools.WeightClassList[body.WeightClass], &QueryCache)
	c.JSON(http.StatusOK, fedData)
}

// Events godoc
//
//		@Summary	Fetch available event metadata within a set date range
//		@Schemes
//		@Description	Metadata shows the name, federation and date of the event along with the filename in the event_data folder.
//		@Tags			OPTIONS Requests
//	 @Param startdate query string false "Start date to filter from"
//	 @Param enddate query string false "End date to filter to"
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	 structs.EventsList
//		@Failure		204	{object}	nil
//		@Router			/events/list [options]
func Events(c *gin.Context) {
	var response structs.EventsList
	var query structs.EventSearch
	if err := c.BindJSON(&query); err != nil {
		abortErr := c.AbortWithError(http.StatusBadRequest, err)
		log.Println(abortErr)
		return
	}

	response.Events = EventsData.FetchEventWithinDate(query.StartDate, query.EndDate)
	c.JSON(http.StatusOK, response)
}

// SingleEvent godoc
//
//		@Summary	Fetch a single event
//		@Schemes
//		@Description	Fetch a single event by ID and federation.
//		@Tags			GET Requests
//	 @Param federation body string true "Federation of the event"
//	 @Param id body string true "ID of the event"
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	 []structs.LeaderboardResponse
//		@Failure		204	{object}	nil
//		@Router			/events [get]
func SingleEvent(c *gin.Context) {
	var response structs.LeaderboardResponse
	var federation, fedExists = c.GetQuery("fed")
	var csvID, idExists = c.GetQuery("id")
	if fedExists && idExists {
		response.Data = dbtools.LoadSingleEvent(federation, csvID)
	} else {
		var eventNameReq, nameExists = c.GetQuery("name")
		if nameExists {
			federation, csvID = EventsData.FetchEventByName(eventNameReq)
			response.Data = dbtools.LoadSingleEvent(federation, csvID)
		}
	}

	response.Size = len(response.Data)
	if response.Size == 0 {
		c.JSON(http.StatusNoContent, nil)
		return
	}
	c.JSON(http.StatusOK, response)
}

// IssueReport godoc
//
//		@Summary	Report an issue with a lift
//		@Schemes
//		@Description	Report an issue with a lift to the discord server
//		@Tags			POST Requests
//	 @Param reportedLift body structs.LiftReport true "Lift to report"
//	 @Param comments body string true "Comments"
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	nil
//		@Router			/issue [post]
func IssueReport(c *gin.Context) {
	var report structs.LiftReport
	if err := c.BindJSON(&report); err != nil {
		abortErr := c.AbortWithError(http.StatusBadRequest, err)
		log.Println(abortErr)
		return
	}
	log.Printf("Issue report received: %s\n", report.Comments)
	_, err := DiscoKaren.PostMessage(report.ReportedLift.DiscordPrint() + "\nReport Comments: *" + report.Comments + "*")
	if err != nil {
		log.Println("Failed to post message to discord")
	}
	c.JSON(http.StatusOK, nil)
}
