package challenges

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UniqueRequest struct {
	Numbers []int `json:"numbers"`
}

// UniqueOrDuplicatesEndpoint godoc
// @Summary Identify unique and duplicate numbers
// @Description Receives a list of integers and separates them into unique and duplicate values.
// @Tags Arrays
// @Accept json
// @Produce json
// @Param input body UniqueRequest true "List of integers"
// @Success 200 {objects} map[string]interface{}
// @Failure 400 {objects} map[string]string
// @Router /unique_or_duplicates [post]
func UniqueOrDuplicatesEndpoint(c *gin.Context) {
	var req UniqueRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	freq := make(map[int]int)
	for _, n := range req.Numbers {
		freq[n]++
	}

	unique := []int{}
	duplicates := []int{}
	for n, count := range freq {
		if count == 1 {
			unique = append(unique, n)
		} else {
			duplicates = append(duplicates, n)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"unique":     unique,
		"duplicates": duplicates,
	})
}
