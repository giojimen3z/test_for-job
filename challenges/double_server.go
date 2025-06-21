package challenges

type In struct {
	Number int32
	Resp   chan int32
}

// Canal global
var ServerChan = make(chan In, 5000)

// Inicializa el servidor en segundo plano
func StartDoubleServer() {
	go func() {
		for req := range ServerChan {
			req.Resp <- req.Number * 2
		}
	}()
}
