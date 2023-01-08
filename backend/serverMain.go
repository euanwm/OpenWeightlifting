package main

import (
	"backend/dbtools"
	"backend/enum"
	"backend/lifter"
	"backend/structs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var processedLeaderboard structs.LeaderboardData

func getTest(c *gin.Context) {
	hour, min, sec := time.Now().Clock()
	retStruct := structs.TestPayload{Hour: hour, Min: min, Sec: sec}
	c.JSON(http.StatusOK, retStruct)
}

func getSearchName(c *gin.Context) {
	if len(c.Query("name")) >= 3 {
		search := structs.NameSearch{NameStr: c.Query("name")}
		suggestions := lifter.NameSearch(search.NameStr, &processedLeaderboard.AllNames)
		results := structs.NameSearchResults{Names: processedLeaderboard.FetchNames(suggestions)}
		c.JSON(http.StatusOK, results)
	}
}

func postLifterRecord(c *gin.Context) {
	lifterSearch := structs.NameSearch{}
	if err := c.BindJSON(&lifterSearch); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
	}
	lifterDetails := lifter.FetchLifts(lifterSearch, &processedLeaderboard)
	lifterDetails.Lifts = dbtools.SortDate(lifterDetails.Lifts)
	if len(lifterDetails.Lifts) != 0 {
		c.JSON(http.StatusOK, lifterDetails)
	} else if len(lifterDetails.Lifts) == 0 {
		c.JSON(http.StatusNoContent, nil)
	}
}

//Main leaderboard function
func postLeaderboard(c *gin.Context) {
	body := structs.LeaderboardPayload{}
	if err := c.BindJSON(&body); err != nil {
		abortErr := c.AbortWithError(http.StatusBadRequest, err)
		log.Println(abortErr)
		return
	}
	//todo: add in query checker, so it doesn't shit the bed on a bad query

	if body != enum.DefaultPayload {
		fedData := dbtools.Filter(processedLeaderboard.Query(body.SortBy, body.Gender), body.Federation, body.WeightClass, body.Start, body.Stop)
		c.JSON(http.StatusOK, fedData)

	} else {
		c.JSON(http.StatusOK, processedLeaderboard.Query(body.SortBy, body.Gender)[body.Start:body.Stop])
	}
}

func buildDatabase() {
	log.Println("buildDatabase called...")
	bigData := dbtools.CollateAll()
	male, female, _ := dbtools.SortGender(bigData) // Throwaway the unknown genders as they're likely really young kids
	const maxSize int = 5000
	leaderboardTotal := &structs.LeaderboardData{
		AllNames:        append(male.ProcessNames(), female.ProcessNames()...),
		AllData:         append(male.Lifts, female.Lifts...),
		MaleTotals:      dbtools.TopPerformance(male.Lifts, enum.Total, maxSize),
		FemaleTotals:    dbtools.TopPerformance(female.Lifts, enum.Total, maxSize),
		MaleSinclairs:   dbtools.TopPerformance(male.Lifts, enum.Sinclair, maxSize),
		FemaleSinclairs: dbtools.TopPerformance(female.Lifts, enum.Sinclair, maxSize),
	}
	processedLeaderboard = *leaderboardTotal
	log.Println("Database READY")
}

func CORSConfig(localEnv bool) cors.Config {
	corsConfig := cors.DefaultConfig()
	if localEnv {
		log.Println("Local mode - Disabling CORS nonsense")
		corsConfig.AllowOrigins = []string{"https://www.openweightlifting.org", "http://localhost:3000"}
	} else {
		corsConfig.AllowOrigins = []string{"https://www.openweightlifting.org"}
	}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

func main() {
	go buildDatabase()
	r := gin.Default()
	//It's not a great solution but it'll work
	if len(os.Args) > 1 && os.Args[1] == "local" {
		r.Use(cors.New(CORSConfig(true)))
	} else {
		r.Use(cors.New(CORSConfig(false)))
	}
	r.GET("test", getTest)
	r.POST("leaderboard", postLeaderboard)
	r.GET("search", getSearchName)
	r.POST("lifter", postLifterRecord)
	err := r.Run()
	if err != nil {
		log.Fatal("Failed to run server")
	}
}
