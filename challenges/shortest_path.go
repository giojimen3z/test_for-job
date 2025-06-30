package challenges

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GridRequest struct {
	Grid [][]int `json:"grid"`
}

// ShortestPathEndpoint godoc
// @Summary Compute the shortest path in a grid (placeholder)
// @Description Accepts a grid and returns a fake response for the shortest path.
// @Tags Matrix
// @Accept json
// @Produce json
// @Param input body GridRequest true "2D grid structure"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Router /shortest_path [post]
func ShortestPathEndpoint(c *gin.Context) {
	var req GridRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid grid"})
		return
	}

	// Dummy response for now
	c.JSON(http.StatusOK, gin.H{"path": "placeholder", "length": 0})
}
