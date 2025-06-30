package challenges

import (
	"net/http"
	"sort"
	"strings"

	"github.com/gin-gonic/gin"
)

func sortString(s string) string {
	r := strings.Split(s, "")
	sort.Strings(r)
	return strings.Join(r, "")
}

// IsAnagramEndpoint godoc
// @Summary Check if two words are anagrams
// @Description Determines whether two words are anagrams, ignoring case.
// @Tags Text
// @Produce json
// @Param word1 path string true "First word"
// @Param word2 path string true "Second word"
// @Success 200 {object} map[string]bool
// @Router /is_anagram/{word1}/{word2} [get]
func IsAnagramEndpoint(c *gin.Context) {
	w1 := c.Param("word1")
	w2 := c.Param("word2")

	if sortString(strings.ToLower(w1)) == sortString(strings.ToLower(w2)) {
		c.JSON(http.StatusOK, gin.H{"anagram": true})
	} else {
		c.JSON(http.StatusOK, gin.H{"anagram": false})
	}
}
