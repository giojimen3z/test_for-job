package challenges

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CountCharactersRequest struct {
	Text string `json:"text"`
}

// CountCharactersEndpoint godoc
// @Summary Count character frequencies in a string
// @Description Returns the frequency of each character in the input string.
// @Tags Text
// @Accept json
// @Produce json
// @Param input body CountCharactersRequest true "Input string"
// @Success 200 {objects} map[string]map[string]int
// @Failure 400 {objects} map[string]string
// @Router /count_characters [post]
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
