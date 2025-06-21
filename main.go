package main

import (
	"github.com/gin-gonic/gin"
	"test_for-job/challenges"
	"test_for-job/routes"
)

func main() {
	router := gin.Default()

	// --- Endpoints: Lógica en challenges ---
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

	// --- Endpoints: Lógica en routes ---
	router.POST("/fork-reader", routes.ForkReaderHandler)
	router.POST("/double-server", routes.DoubleServerHandler)
	router.POST("/pair-sum", routes.PairSumHandler)
	router.POST("/browser-navigator", routes.BrowserNavigatorHandler)
	router.POST("/channel-pair-sum", routes.ChannelPairSumHandler)

	// --- Inicialización de servidores internos ---
	challenges.StartDoubleServer()

	// --- Iniciar servidor HTTP ---
	router.Run(":8080") // Escucha en http://localhost:8080
}
