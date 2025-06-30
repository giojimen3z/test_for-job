package routes

import (
	"net/http"

	"test_for-job/challenges"

	"github.com/gin-gonic/gin"
)

type BrowserNavigatorRequest struct {
	Commands []string `json:"commands"`
}

type BrowserNavigatorResponse struct {
	Log []string `json:"log"`
}

// BrowserNavigatorEndpoint godoc
// @Summary Simulate browser navigation
// @Description Simulates navigating through browser commands like hopTo, backtrack, leapForward
// @Tags Navigation
// @Accept json
// @Produce json
// @Param input body map[string][]string true "List of navigation commands"
// @Success 200 {array} string
// @Router /browser-navigator [post]
func BrowserNavigatorHandler(c *gin.Context) {
	var req BrowserNavigatorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := challenges.SimulateBrowserNavigation(req.Commands)
	c.JSON(http.StatusOK, BrowserNavigatorResponse{Log: result})
}
