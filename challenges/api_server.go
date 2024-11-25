package challenges

import (
	"github.com/gin-gonic/gin"
)

func APIServerExampleEndpoint(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello, this is an example API endpoint!",
	})
}
