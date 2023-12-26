package main

import (
	"backend/dbtools"
	docs "github.com/euanwm/OpenWeightlifting/backend/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
)

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

func setupCORS(r *gin.Engine) {
	// It's not a great solution but it'll work
	if len(os.Args) > 1 && os.Args[1] == "local" {
		r.Use(cors.New(CORSConfig(true)))
	} else {
		r.Use(cors.New(CORSConfig(false)))
	}
}

func buildServer() *gin.Engine {
	log.Println("Starting server...")
	dbtools.BuildDatabase(&processedLeaderboard)
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	setupCORS(r)
	r.GET("time", ServerTime)
	r.POST("leaderboard", Leaderboard)
	r.GET("search", SearchName)
	r.POST("lifter", LifterRecord)
	r.POST("history", LifterHistory)
	r.POST("event", EventResult)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

// CacheMeOutsideHowBoutDat - Precaches data on startup on a separate thread due to container timeout constraints.
func CacheMeOutsideHowBoutDat() {
	log.Println("Precaching data...")
	for n, query := range dbtools.PreCacheQuery {
		log.Println("Caching query: ", n)
		_, _ = QueryCache.CheckQuery(query)
		liftdata := processedLeaderboard.Select(query.SortBy)
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
	apiServer := buildServer()
	go CacheMeOutsideHowBoutDat()
	err := apiServer.Run()
	if err != nil {
		log.Fatal("Failed to run server")
	}
}
