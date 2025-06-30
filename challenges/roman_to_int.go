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

// RomanToIntEndpoint godoc
// @Summary Convert Roman numeral to integer
// @Description Converts a Roman numeral string to its integer representation.
// @Tags Conversion
// @Produce json
// @Param roman path string true "Roman numeral (e.g., X, IX, MCMXCIV)"
// @Success 200 {objects} map[string]interface{}
// @Failure 400 {objects} map[string]string
// @Router /roman_to_int/{roman} [get]

func RomanToIntEndpoint(c *gin.Context) {
	roman := c.Param("roman")
	integer := romanToInt(roman)

	c.JSON(200, gin.H{
		"roman":   roman,
		"integer": integer,
	})
}
