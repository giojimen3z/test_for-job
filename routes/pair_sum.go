package routes

import (
	"net/http"

	"test_for-job/challenges"

	"github.com/gin-gonic/gin"
)

type PairSumRequest struct {
	Ascending  []int `json:"ascending"`
	Descending []int `json:"descending"`
	Target     int   `json:"target"`
}

type PairSumResponse struct {
	Count int `json:"pairs"`
}

func PairSumHandler(c *gin.Context) {
	var req PairSumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count := challenges.CheckPairSum(req.Ascending, req.Descending, req.Target)
	c.JSON(http.StatusOK, PairSumResponse{Count: count})
}
