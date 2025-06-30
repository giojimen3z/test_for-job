package challenges

import (
	"net/http"
	"time"

	"github.com/andelf/go-curl"
	"github.com/gin-gonic/gin"
)

type URLRequest struct {
	URLs []string `json:"urls"`
}

// GetTimeForUrlsEndpoint godoc
// @Summary Measure response time for multiple URLs
// @Description Sends HTTP requests to a list of URLs and returns the response time in seconds for each.
// @Tags Network
// @Accept json
// @Produce json
// @Param input body URLRequest true "List of URLs"
// @Success 200 {objects} map[string]float64
// @Failure 400 {objects} map[string]string
// @Router /get_time_for_urls [post]
func GetTimeForUrlsEndpoint(c *gin.Context) {
	var req URLRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.URLs) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	results := make(map[string]float64)
	curlInit := curl.EasyInit()
	defer curlInit.Cleanup()

	for _, urlStr := range req.URLs {
		curlInit.Setopt(curl.OPT_URL, urlStr)
		curlInit.Setopt(curl.OPT_WRITEFUNCTION, func(buf []byte, userdata interface{}) bool { return true })

		start := time.Now()
		_ = curlInit.Perform()
		duration := time.Since(start)
		res, _ := curlInit.Getinfo(curl.INFO_RESPONSE_CODE)
		if resAsInt, ok := res.(int); ok && resAsInt < 400 {
			results[urlStr] = duration.Seconds()
		} else {
			results[urlStr] = -1
		}
	}

	c.JSON(http.StatusOK, gin.H{"results": results})
}
