package challenges

import (
	"sync"

	"github.com/gin-gonic/gin"
)

type Stack struct {
	data []int
	mu   sync.Mutex
}

func (s *Stack) Push(value int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, value)
}

func (s *Stack) Pop() (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.data) == 0 {
		return 0, false
	}
	value := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return value, true
}

// Handler para el endpoint de Stack
func StackEndpoint(c *gin.Context) {
	stack := &Stack{data: []int{}}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	results := []int{}
	for i := 0; i < 3; i++ {
		if value, ok := stack.Pop(); ok {
			results = append(results, value)
		}
	}

	c.JSON(200, gin.H{
		"message": "Stack operations completed successfully",
		"results": results,
	})
}
