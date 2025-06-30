package challenges

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// MultiplesSumEndpoint godoc
// @Summary Sum multiples of 3 or 5 below a limit
// @Description Returns the sum of all numbers below the given limit that are divisible by 3 or 5.
// @Tags Math
// @Produce json
// @Param limit path int true "Upper limit (exclusive)"
// @Success 200 {objects} map[string]int
// @Failure 400 {objects} map[string]string
// @Router /multiples_sum/{limit} [get]
func MultiplesSumEndpoint(c *gin.Context) {
	limitStr := c.Param("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
		return
	}

	sum := 0
	for i := 1; i < limit; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	c.JSON(http.StatusOK, gin.H{"sum": sum})
}
