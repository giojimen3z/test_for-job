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

func BrowserNavigatorHandler(c *gin.Context) {
	var req BrowserNavigatorRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := challenges.SimulateBrowserNavigation(req.Commands)
	c.JSON(http.StatusOK, BrowserNavigatorResponse{Log: result})
}
