package challenges

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PairSumRequest struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

func PairSumCountEndpoint(c *gin.Context) {
	var req PairSumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	seen := make(map[int]bool)
	count := 0
	for _, num := range req.Numbers {
		if seen[req.Target-num] {
			count++
		}
		seen[num] = true
	}

	c.JSON(http.StatusOK, gin.H{"count": count})
}
