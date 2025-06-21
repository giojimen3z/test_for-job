package routes

import (
	"net/http"

	"test_for-job/challenges"

	"github.com/gin-gonic/gin"
)

type ChannelPairSumRequest struct {
	Ascending  []int `json:"ascending"`
	Descending []int `json:"descending"`
	Target     int64 `json:"target"`
}

type ChannelPairSumResponse struct {
	Pairs uint32 `json:"pairs"`
}

func ChannelPairSumHandler(c *gin.Context) {
	var req ChannelPairSumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	size := len(req.Ascending)
	if size != len(req.Descending) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ascending and descending must be the same length"})
		return
	}

	ascChan := make(chan int, size)
	descChan := make(chan int, size)

	for _, v := range req.Ascending {
		ascChan <- v
	}
	for _, v := range req.Descending {
		descChan <- v
	}

	close(ascChan)
	close(descChan)

	count := challenges.ChannelPairSum(ascChan, descChan, req.Target, size)
	c.JSON(http.StatusOK, ChannelPairSumResponse{Pairs: count})
}
