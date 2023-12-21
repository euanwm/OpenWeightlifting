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

var QueryCache dbtools.QueryCache

// PingExample godoc
//
//	@Summary	Checking the servers localtime
//	@Schemes
//	@Description	Returns the current server time.
//	@Tags			Utilities and Testing
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.ContainerTime
//	@Router			/time [get]
func ServerTime(c *gin.Context) {
	hour, mins, sec := time.Now().Clock()
	retStruct := structs.ContainerTime{Hour: hour, Min: mins, Sec: sec}
	c.JSON(http.StatusOK, retStruct)
}

// PingExample godoc
//
//	@Summary	how to use the name search endpoint
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.NameSearchResults
//	@Router			/search [get]
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

// PingExample godoc
//
//	@Summary	how to use the event result endpoint
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{slice}	 []structs.Entry
//	@Router			/event [post]
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

// PingExample godoc
//
//	@Summary	how to use the lifter record endpoint
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.ChartData
//	@Router			/lifter [post]
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

// PingExample godoc
//
//	@Summary	how to use the lifter history endpoint
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.LifterHistory
//	@Router			/history [post]
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

// PingExample godoc
//
//	@Summary	how to use the leaderboard endpoint
//	@Schemes
//	@Description	do ping
//	@Tags			example
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	structs.LeaderboardResponse
//	@Router			/leaderboard [post]
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
