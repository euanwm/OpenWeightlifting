package main

import (
	"backend/dbtools"
	_ "backend/docs"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CORSConfig() cors.Config {
	corsConfig := cors.DefaultConfig()
	if os.Getenv("GIN_MODE") != gin.ReleaseMode {
		log.Println("Local mode - Disabling CORS nonsense")
		corsConfig.AllowOrigins = []string{"https://www.openweightlifting.org", "https://openweightlifting.org", "http://localhost:3000", "http://frontend-app:3000", "*"}
	} else {
		corsConfig.AllowOrigins = []string{"https://www.openweightlifting.org", "https://openweightlifting.org", "https://owl-v2-production.up.railway.app", "https://alpha.openweightlifting.org"}
	}
	corsConfig.AllowCredentials = true
	corsConfig.AddAllowHeaders("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers", "Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "X-Requested-With", "Authorization")
	corsConfig.AddAllowMethods("GET", "POST", "PUT", "DELETE")
	return corsConfig
}

func buildServer() *gin.Engine {
	log.Println("Starting server...")
	dbtools.BuildDatabase(&LeaderboardData, &EventsData, &LifterRoster)
	r := gin.Default()
	r.Use(cors.New(CORSConfig()))
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.GET("time", ServerTime)
	r.GET("leaderboard", Leaderboard)
	r.POST("leaderboard/search", LeaderboardSearch)
	r.GET("search", SearchName)
	r.GET("search/similarity", SimilarNameSearch)
	r.GET("graph", LifterGraph)
	r.GET("history", LifterHistory)
	r.GET("events/list", Events)
	r.GET("events", SingleEvent)
	r.GET("rivals", Rival)
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}

// CacheMeOutsideHowBoutDat - Precaches data on startup on a separate thread due to container timeout constraints.
func CacheMeOutsideHowBoutDat() {
	log.Println("Precaching data...")
	for n, query := range dbtools.PreCacheQuery() {
		log.Println("Caching query: ", n)
		_, _ = QueryCache.CheckQuery(query)
		liftdata := LeaderboardData.Select(query.SortBy)
		dbtools.PreCacheFilter(liftdata, query, dbtools.WeightClassList[query.WeightClass], &QueryCache)
	}
	log.Println("Caching complete")

	if os.Getenv("GIN_MODE") == gin.ReleaseMode {
		log.Println("Syncíng frontend")
		err := backendAlive()
		if err != nil {
			log.Println("Error notifying frontend: ", err)
		}
	}
}

func backendAlive() error {
	version := os.Getenv("RELEASE_VERSION")
	endpoint := os.Getenv("FRONTEND_ENDPOINT")

	if endpoint == "" {
		return fmt.Errorf("FRONTEND_ENDPOINT environment variable not set")
	}

	if version == "" {
		return fmt.Errorf("RELEASE_VERSION environment variable not set")
	}

	log.Println("Pushing version to frontend: ", version)
	jsonBody := []byte(fmt.Sprintf(`{"version": "%s"}`, version))
	bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest(http.MethodPut, endpoint, bodyReader)
	if err != nil {
		log.Println("big err: ", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	log.Println("Response Status:", resp.Status)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil

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
	err := apiServer.Run() // listen and serve on
	if err != nil {
		log.Fatal("Failed to run server")
	}
}
