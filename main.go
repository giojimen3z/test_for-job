package main

import (
	_ "net/http/pprof"

	_ "test_for-job/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"test_for-job/challenges"
	"test_for-job/routes"
)

// @title Golang Code Challenges API
// @version 1.0
// @description This is a collection of code challenge endpoints written in Go using Gin.
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	// --- Endpoints: Challenge logic ---
	router.GET("/api", challenges.APIServerExampleEndpoint)
	router.GET("/stack", challenges.StackEndpoint)
	router.GET("/concurrency_sim", challenges.ConcurrencySimEndpoint)
	router.GET("/codereview", challenges.CodeReviewEndpoint)
	router.GET("/roman_to_int/:roman", challenges.RomanToIntEndpoint)
	router.GET("/multiples_sum/:limit", challenges.MultiplesSumEndpoint)
	router.GET("/is_anagram/:word1/:word2", challenges.IsAnagramEndpoint)

	router.POST("/count_characters", challenges.CountCharactersEndpoint)
	router.POST("/get_time_for_urls", challenges.GetTimeForUrlsEndpoint)
	router.POST("/pair_sum_count", challenges.PairSumCountEndpoint)
	router.POST("/shortest_path", challenges.ShortestPathEndpoint)
	router.POST("/unique_or_duplicates", challenges.UniqueOrDuplicatesEndpoint)
	router.POST("/conncache_scaffold", challenges.ConnCacheEndpoint)

	// --- Endpoints: Routes logic ---
	router.POST("/fork-reader", routes.ForkReaderHandler)
	router.POST("/double-server", routes.DoubleServerHandler)
	router.POST("/pair-sum", routes.PairSumHandler)
	router.POST("/browser-navigator", routes.BrowserNavigatorHandler)
	router.POST("/channel-pair-sum", routes.ChannelPairSumHandler)

	// --- Internal server initialization ---
	challenges.StartDoubleServer()

	// --- Swagger UI route ---
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// --- Start the HTTP server ---
	router.Run(":8080")
}
