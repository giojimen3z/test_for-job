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

// DoubleServerHandler godoc
// @Summary Double numbers using an internal channel-based server
// @Description Sends a list of integers to a background server via a channel and returns each number multiplied by two.
// @Tags Concurrency
// @Accept json
// @Produce json
// @Param input body DoubleServerRequest true "List of numbers to double"
// @Success 200 {objects} DoubleServerResponse
// @Failure 400 {objects} map[string]string
// @Router /double-server [post]
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
