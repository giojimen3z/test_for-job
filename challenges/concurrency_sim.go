package challenges

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func ConcurrencySimEndpoint(c *gin.Context) {
	var wg sync.WaitGroup
	results := make([]string, 5)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(i) * 100 * time.Millisecond)
			results[i] = fmt.Sprintf("Task %d done", i)
		}(i)
	}

	wg.Wait()
	c.JSON(http.StatusOK, gin.H{"results": results})
}
