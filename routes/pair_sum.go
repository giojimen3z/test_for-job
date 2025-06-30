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

// PairSumHandler godoc
// @Summary Count valid pairs that sum to target
// @Description Given two arrays (one ascending, one descending) and a target value, returns the number of pairs (one from each array) that sum to the target.
// @Tags Array
// @Accept json
// @Produce json
// @Param input body PairSumRequest true "Two arrays and target value"
// @Success 200 {objects} PairSumResponse
// @Failure 400 {objects} map[string]string
// @Router /pair-sum [post]
func PairSumHandler(c *gin.Context) {
	var req PairSumRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	count := challenges.CheckPairSum(req.Ascending, req.Descending, req.Target)
	c.JSON(http.StatusOK, PairSumResponse{Count: count})
}
