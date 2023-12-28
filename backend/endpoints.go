package main //nolint:typecheck

import (
	"backend/dbtools"
	"backend/enum"
	"backend/events"
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

var processedLeaderboard structs.LeaderboardData

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
//	 @Param name query string true "name"
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.NameSearchResults
//		@Router			/search [get]
func SearchName(c *gin.Context) {
	const maxResults = 50
	if len(c.Query("name")) >= 3 {
		search := structs.NameSearch{NameStr: c.Query("name")}
		results := structs.NameSearchResults{Names: lifter.NameSearch(search.NameStr, &processedLeaderboard.AllTotals)}
		// todo: remove this and implement a proper solution
		if len(results.Names) > maxResults {
			results.Names = results.Names[:maxResults]
		}
		c.JSON(http.StatusOK, results)
	}
}

// EventResult godoc
//
//		@Summary	Pull a specific event by name
//		@Schemes
//		@Description	Requires a case-sensitive event name to be passed to it. This is still a work in progress.
//		@Tags			POST Requests
//	 @Param name body string true "name"
//		@Accept			json
//		@Produce		json
//		@Success		200	{array}	 []structs.Entry
//		@Failure		204	{object}	nil
//		@Router			/event [post]
func EventResult(c *gin.Context) {
	eventSearch := structs.NameSearch{}
	if err := c.BindJSON(&eventSearch); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
	}
	eventData := events.FetchEvent(eventSearch.NameStr, &processedLeaderboard)
	if len(eventData) != 0 {
		c.JSON(http.StatusOK, eventData)
	} else {
		c.JSON(http.StatusNoContent, nil)
	}
}

// LifterRecord godoc
//
//		@Summary	Retrieve a lifter's record for use with ChartJS
//		@Schemes
//		@Description	This is used within the lifter page to display a lifter's record. It returns a JSON object that can be used with ChartJS without having to do any additional processing.
//		@Tags			POST Requests
//	 @Param name body string true "name"
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.ChartData
//	 @Failure		204	{object}	nil
//		@Router			/lifter [post]
func LifterRecord(c *gin.Context) {
	lifterSearch := structs.NameSearch{}
	if err := c.BindJSON(&lifterSearch); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
	}

	lifterDetails := lifter.FetchLifts(lifterSearch, &processedLeaderboard)
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
//		@Tags			POST Requests
//	 @Param name body string true "name"
//		@Accept			json
//		@Produce		json
//		@Success		200	{object}	structs.LifterHistory
//	 @Failure		204	{object}	nil
//		@Router			/history [post]
func LifterHistory(c *gin.Context) {
	lifterSearch := structs.NameSearch{}
	if err := c.BindJSON(&lifterSearch); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
	}

	lifterDetails := lifter.FetchLifts(lifterSearch, &processedLeaderboard)
	lifterDetails.Lifts = dbtools.SortDate(lifterDetails.Lifts)
	lifterDetails.Graph = lifterDetails.GenerateChartData()
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
//		@Tags			POST Requests
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
	body := structs.LeaderboardPayload{}
	if err := c.BindJSON(&body); err != nil {
		abortErr := c.AbortWithError(http.StatusBadRequest, err)
		log.Println(abortErr)
		return
	}

	// todo: remove this once the frontend filters have been updated to suit
	switch body.Year {
	case 69:
		body.StartDate = enum.ZeroDate
		body.EndDate = enum.MaxDate
	default:
		body.StartDate = strconv.Itoa(body.Year) + "-01-01"
		body.EndDate = strconv.Itoa(body.Year+1) + "-01-01"
	}

	leaderboardData := processedLeaderboard.Select(body.SortBy) // Selects either total or sinclair sorted leaderboard
	fedData := dbtools.FilterLifts(*leaderboardData, body, dbtools.WeightClassList[body.WeightClass], &QueryCache)
	c.JSON(http.StatusOK, fedData)
}
