package challenges

type In struct {
	Number int32
	Resp   chan int32
}

// Global channel used for communication with the background server
var ServerChan = make(chan In, 5000)

// StartDoubleServer launches a background goroutine that listens on ServerChan.
// For each incoming request, it doubles the input number and sends it back through the response channel.
func StartDoubleServer() {
	go func() {
		for req := range ServerChan {
			req.Resp <- req.Number * 2
		}
	}()
}
