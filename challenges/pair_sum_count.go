package challenges

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type PairSumRequest struct {
	Numbers []int `json:"numbers"`
	Target  int   `json:"target"`
}

// PairSumCountEndpoint godoc
// @Summary Count pairs that sum to a target
// @Description Returns how many unique pairs in an array sum up to a given target value.
// @Tags Array
// @Accept json
// @Produce json
// @Param input body PairSumRequest true "Array of numbers and target value"
// @Success 200 {objects} map[string]int
// @Failure 400 {objects} map[string]string
// @Router /pair_sum_count [post]
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
