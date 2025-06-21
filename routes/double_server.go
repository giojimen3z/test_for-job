package routes

import (
	"net/http"

	"test_for-job/challenges"

	"github.com/gin-gonic/gin"
)

type DoubleServerRequest struct {
	Numbers []int32 `json:"numbers"`
}

type DoubleServerResponse struct {
	Results []int32 `json:"results"`
}

func DoubleServerHandler(c *gin.Context) {
	var req DoubleServerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	results := make([]int32, 0, len(req.Numbers))
	for _, num := range req.Numbers {
		respChan := make(chan int32)
		challenges.ServerChan <- challenges.In{
			Number: num,
			Resp:   respChan,
		}
		results = append(results, <-respChan)
	}

	c.JSON(http.StatusOK, DoubleServerResponse{Results: results})
}
