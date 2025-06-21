package challenges

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GridRequest struct {
	Grid [][]int `json:"grid"`
}

func ShortestPathEndpoint(c *gin.Context) {
	var req GridRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid grid"})
		return
	}

	// Dummy response for now
	c.JSON(http.StatusOK, gin.H{"path": "placeholder", "length": 0})
}
