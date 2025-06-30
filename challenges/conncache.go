package challenges

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

type Conn struct {
	Addr string `json:"addr"`
}

func NewConn(addr string) *Conn {
	return &Conn{Addr: addr}
}

func (conn *Conn) Dial() error {
	return nil
}

func (conn *Conn) Close() error {
	return nil
}

type Cache struct {
	mu    sync.Mutex
	store map[string]*Conn
}

func NewCache() *Cache {
	return &Cache{store: make(map[string]*Conn)}
}

func (cache *Cache) Get(addr string) (*Conn, error) {
	cache.mu.Lock()
	defer cache.mu.Unlock()

	if conn, ok := cache.store[addr]; ok {
		return conn, nil
	}

	newConn := NewConn(addr)
	err := newConn.Dial()
	if err != nil {
		return nil, err
	}

	cache.store[addr] = newConn
	return newConn, nil
}

func (cache *Cache) Cleanup() {
	// not implemented yet
}

type ConnCacheRequest struct {
	Sequence []string `json:"sequence"`
}

type ConnCacheResult struct {
	Successful bool   `json:"successful"`
	Message    string `json:"message"`
}

// ConnCacheEndpoint godoc
// @Summary Simulate and validate cached connections
// @Description Accepts a sequence of addresses, simulates cached connection reuse, and validates correct behavior.
// @Tags Cache
// @Accept json
// @Produce json
// @Param input body ConnCacheRequest true "Sequence of connection addresses"
// @Success 200 {objects} ConnCacheResult
// @Failure 400 {objects} map[string]string
// @Failure 500 {objects} map[string]string
// @Router /conncache_scaffold [post]
func ConnCacheEndpoint(c *gin.Context) {
	var req ConnCacheRequest
	if err := c.ShouldBindJSON(&req); err != nil || len(req.Sequence) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input sequence"})
		return
	}

	cache := NewCache()
	conns := make(map[string]*Conn)

	for _, addr := range req.Sequence {
		conn, err := cache.Get(addr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("connection failed for %s", addr)})
			return
		}
		conns[addr] = conn
	}

	// Comparaciones similares al main
	first := conns[req.Sequence[0]]
	last := conns[req.Sequence[len(req.Sequence)-1]]

	uniqueAddrs := make(map[*Conn]struct{})
	for _, conn := range conns {
		uniqueAddrs[conn] = struct{}{}
	}

	// Resultado de prueba
	if len(uniqueAddrs) != len(conns) && first == last {
		c.JSON(http.StatusOK, ConnCacheResult{
			Successful: false,
			Message:    "Error: same Conn used for different addresses",
		})
		return
	}

	c.JSON(http.StatusOK, ConnCacheResult{
		Successful: true,
		Message:    "Successful completion",
	})
}
