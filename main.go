package main

import (
	"github.com/gin-gonic/gin"
	"test_for-job/challenges"
)

func main() {
	router := gin.Default()

	// Rutas para los desafíos
	router.GET("/stack", challenges.StackEndpoint)
	router.GET("/roman_to_int/:roman", challenges.RomanToIntEndpoint)
	router.GET("/api", challenges.APIServerExampleEndpoint)

	// Iniciar el servidor
	router.Run(":8080") // Escucha en http://localhost:8080
}
