package challenges

import (
	"github.com/gin-gonic/gin"
)

// APIServerExampleEndpoint godoc
// @Summary Example API endpoint
// @Description Returns a basic greeting message.
// @Tags Example
// @Produce json
// @Success 200 {objects} map[string]string
// @Router /api [get]
func APIServerExampleEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, this is an example API endpoint!",
	})
}
