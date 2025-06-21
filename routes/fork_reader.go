package routes

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"test_for-job/challenges"
)

type ForkReaderRequest struct {
	Data string `json:"data"`
}

type ForkReaderResponse struct {
	W1 string `json:"w1"`
	W2 string `json:"w2"`
}

func ForkReaderHandler(c *gin.Context) {
	var req ForkReaderRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r := bytes.NewReader([]byte(req.Data))
	var w1, w2 bytes.Buffer

	err := challenges.ForkReader(r, &w1, &w2)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ForkReaderResponse{
		W1: w1.String(),
		W2: w2.String(),
	})
}
