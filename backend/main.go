package main

import (
	"backend/dbtools"
	"backend/discordbot"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	if os.Getenv("GIN_MODE") != gin.ReleaseMode {
		log.Println("Local mode - Disabling CORS nonsense")
		corsConfig.AllowOrigins = []string{"https://www.openweightlifting.org", "http://localhost:3000", "http://frontend-app:3000", "*"}
	} else {
		corsConfig.AllowOrigins = []string{"https://www.openweightlifting.org"}
	}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

func setupDiscordBot(bot *discordbot.DiscordBot) {
	token := os.Getenv("DISCORD_TOKEN")
	*bot, _ = discordbot.New(token)
	err := bot.OpenConnection()
	if err != nil {
		log.Println("Failed to open discord connection")
	}
	bot.Channel = os.Getenv("ISSUES_CHANNEL")
	if err != nil {
		log.Println("Failed to post message to discord")
	}
	log.Println("Discord bot started")
}

func buildServer() *gin.Engine {
	log.Println("Starting server...")
	dbtools.BuildDatabase(&LeaderboardData, &EventsData)
	r := gin.Default()
	r.Use(cors.New(CORSConfig()))
	r.GET("time", ServerTime)
	r.GET("leaderboard", Leaderboard)
	r.GET("search", SearchName)
	r.GET("graph", LifterGraph)
	r.GET("history", LifterHistory)
	r.POST("events/list", Events)
	r.GET("events", SingleEvent)
	r.POST("issue", IssueReport)
	return r
}

// CacheMeOutsideHowBoutDat - Precaches data on startup on a separate thread due to container timeout constraints.
func CacheMeOutsideHowBoutDat() {
	log.Println("Precaching data...")
	for n, query := range dbtools.PreCacheQuery() {
		log.Println("Caching query: ", n)
		_, _ = QueryCache.CheckQuery(query)
		liftdata := LeaderboardData.Select(query.SortBy)
		dbtools.PreCacheFilter(*liftdata, query, dbtools.WeightClassList[query.WeightClass], &QueryCache)
	}
	log.Println("Caching complete")
}

// @title OpenWeightlifting API
// @description This is the API for OpenWeightlifting.org
// @BasePath /
// @version 1.0
// @contact.name Euan Meston
// @contact.email euan@openweightlifting.org
// @host api.openweightlifting.org
// @schemes https
func main() {
	setupDiscordBot(&DiscoKaren)
	apiServer := buildServer()
	//go CacheMeOutsideHowBoutDat()
	err := apiServer.Run() // listen and serve on
	if err != nil {
		log.Fatal("Failed to run server")
	}
}
