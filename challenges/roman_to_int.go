package challenges

import (
	"github.com/gin-gonic/gin"
)

func romanToInt(roman string) int {
	romanToValue := map[byte]int{
		'I': 1, 'V': 5, 'X': 10,
		'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	result := 0
	for i := 0; i < len(roman); i++ {
		if i < len(roman)-1 && romanToValue[roman[i]] < romanToValue[roman[i+1]] {
			result -= romanToValue[roman[i]]
		} else {
			result += romanToValue[roman[i]]
		}
	}
	return result
}

// Handler para el endpoint de Romanos a Enteros
func RomanToIntEndpoint(c *gin.Context) {
	roman := c.Param("roman")
	integer := romanToInt(roman)

	c.JSON(200, gin.H{
		"roman":   roman,
		"integer": integer,
	})
}
