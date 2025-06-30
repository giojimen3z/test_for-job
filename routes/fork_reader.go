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

// ForkReaderHandler godoc
// @Summary Fork a stream alternately into two outputs
// @Description Reads a string and alternates each byte into two outputs (w1 and w2).
// @Tags Stream
// @Accept json
// @Produce json
// @Param input body ForkReaderRequest true "Input stream data"
// @Success 200 {objects} ForkReaderResponse
// @Failure 400 {objects} map[string]string
// @Failure 500 {objects} map[string]string
// @Router /fork-reader [post]
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
