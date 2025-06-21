package challenges

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UniqueRequest struct {
	Numbers []int `json:"numbers"`
}

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
