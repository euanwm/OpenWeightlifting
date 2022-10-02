package main

import (
	"backend/dbtools"
	"backend/enum"
	"backend/structs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var processedLeaderboard = buildDatabase()

//Main leaderboard function
func postLeaderboard(c *gin.Context) {
	log.Println("postLeaderboard called...")
	body := structs.LeaderboardPayload{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	switch body.Gender {
	case enum.Male:
		c.JSON(http.StatusOK, processedLeaderboard.MaleTotals[body.Start:body.Stop])
	case enum.Female:
		c.JSON(http.StatusOK, processedLeaderboard.FemaleTotals[body.Start:body.Stop])
	default:
		log.Println("Some cunts being wild with it...")
	}
}

func postSinclairs(c *gin.Context) {
	log.Println("postSinclair called...")
	body := structs.LeaderboardPayload{}
	if err := c.BindJSON(&body); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	switch body.Gender {
	case enum.Male:
		c.JSON(http.StatusOK, processedLeaderboard.MaleSinclairs[body.Start:body.Stop])
	case enum.Female:
		c.JSON(http.StatusOK, processedLeaderboard.FemaleSinclairs[body.Start:body.Stop])
	default:
		log.Println("Some cunts being wild with it...")
	}
}

func buildDatabase() (leaderboardTotal *structs.LeaderboardData) {
	log.Println("buildDatabase called...")
	bigData := dbtools.CollateAll()
	male, female, _ := dbtools.SortGender(bigData) // Throwaway the unknown genders as they're likely really young kids
	leaderboardTotal = &structs.LeaderboardData{
		MaleTotals:      dbtools.TopPerformance(male, enum.Total, enum.Male),
		FemaleTotals:    dbtools.TopPerformance(female, enum.Total, enum.Female),
		MaleSinclairs:   dbtools.TopPerformance(male, enum.Sinclair, enum.Male),
		FemaleSinclairs: dbtools.TopPerformance(female, enum.Sinclair, enum.Female),
	}
	return leaderboardTotal
}

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"https://www.openweightlifting.org"}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

func main() {
	r := gin.Default()
	r.Use(cors.New(CORSConfig()))
	r.POST("leaderboard", postLeaderboard)
	r.POST("sinclair", postSinclairs)
	err := r.Run()
	if err != nil {
		log.Fatal("Failed to run server")
	}
}
