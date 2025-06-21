package challenges

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
