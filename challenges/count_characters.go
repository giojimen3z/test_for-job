package challenges

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CountCharactersRequest struct {
	Text string `json:"text"`
}

func CountCharactersEndpoint(c *gin.Context) {
	var req CountCharactersRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	freq := make(map[string]int)
	for _, ch := range strings.Split(req.Text, "") {
		freq[ch]++
	}

	c.JSON(http.StatusOK, gin.H{"frequencies": freq})
}
