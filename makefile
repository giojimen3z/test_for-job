APP_NAME=test-for-job
PORT=8080
PPROF_PORT=6060

.PHONY: run build profile clean

## Build the application
build:
	go build -o $(APP_NAME) main.go

## Run the API with pprof enabled on port 6060
run:
	go run main.go

## Capture a 30-second CPU profile via pprof
profile:
	go tool pprof http://localhost:$(PPROF_PORT)/debug/pprof/profile?seconds=30

## List active goroutines
goroutines:
	curl http://localhost:$(PPROF_PORT)/debug/pprof/goroutine?debug=2

## Clean generated binary
clean:
	rm -f $(APP_NAME)

## Show available make targets
help:
	@echo "Make targets:"
	@echo "  build      Compile the project"
	@echo "  run        Run with pprof enabled on port $(PPROF_PORT)"
	@echo "  profile    Capture 30-second CPU profile"
	@echo "  goroutines Show active goroutines"
	@echo "  clean      Delete the binary"
