# Code Challenges in Go 🚀

## 📖 Description
This project brings together multiple technical challenges implemented in **Go**, exposed through a REST API using the **Gin Gonic** framework. Each challenge addresses a common problem and solves it using idiomatic, safe, and efficient Go code. The project is designed to be easily extensible and testable.

---

## 💡 What does this project include?

1. **Stack Operations**: A concurrency-safe stack implementation for pushing and popping elements.
2. **Roman to Integer Conversion**: Converts Roman numerals to their decimal representation.
3. **Example API**: A basic endpoint that returns a simple greeting.
4. **Sum of Multiples**: Calculates the sum of multiples of 3 or 5 below a given limit.
5. **Pair Sum Count**: Counts how many pairs in an array sum to a target value.
6. **Anagram Check**: Verifies whether two words are anagrams.
7. **Shortest Path in Matrix**: Placeholder for pathfinding in a 2D matrix.
8. **Unique and Duplicate Identifier**: Separates unique and duplicate values in a list.
9. **Concurrency Simulation**: Runs concurrent tasks using goroutines and `sync.WaitGroup`.
10. **URL Response Times**: Measures the latency of multiple URLs using `libcurl`.
11. **Character Counter**: Returns the frequency of each character in a string.
12. **Fork Reader** (`POST /fork-reader`):
    Splits a stream of text alternately byte by byte into two separate outputs (`w1` and `w2`).
13. **Double Server** (`POST /double-server`):
    Accepts an array of numbers, sends them to an internal server via channels, and returns each number doubled.
14. **Pair Sum (Two-pointer)** (`POST /pair-sum`):
    Given two sorted arrays and a target sum, counts how many cross-pairs add up to the target.
15. **Channel Pair Sum** (`POST /channel-pair-sum`):
    Similar to above but uses channels to simulate input and process matching pairs concurrently.
16. **Browser Navigator** (`POST /browser-navigator`):
    Simulates browser navigation using commands like `hopTo`, `backtrack`, and `leapForward`, maintaining back and forward history.

---

## 📦 Bruno API Client Collection

A `collections/` folder includes `.bru` files compatible with [Bruno API Client](https://www.usebruno.com/) for testing all endpoints.

To use:
1. Open Bruno.
2. Import the `collections/` folder.
3. Run the requests directly.

---

## 🧾 Swagger / OpenAPI Documentation

Swagger documentation is available and auto-generated using [Swaggo](https://github.com/swaggo/swag).

To view it:
1. Run the server.
2. Navigate to `http://localhost:8080/swagger/index.html`

To regenerate the Swagger docs:
```bash
swag init --parseDependency --parseInternal
```

---

## 🔬 Profiling Support (`pprof`)

This project supports Go profiling via `net/http/pprof`.

You can profile:
- CPU
- Memory
- Goroutines
- Locks (mutex contention)

### Quick usage:
```bash
make run        # start API with pprof at :6060
make profile    # capture 30-second CPU profile
make goroutines # list active goroutines
```

More details in [`/profiling/README.md`](./profiling/README.md)

---

## 🛠️ Prerequisites

1. **Go**:
    - Download and install Go from [https://go.dev/dl/](https://go.dev/dl/).
    - Verify the installation:
      ```bash
      go version
      ```

2. **Gin Gonic**:
    - The framework will be installed automatically via `go mod tidy`.

---

## 🚀 How to Run This Project

1. **Clone this repository**:
   ```bash
   git clone https://github.com/giojimen3z/test-for-job.git
   cd test-for-job
   ```

2. **Install dependencies**:
   ```bash
   go mod tidy
   ```

3. **Run the server**:
   ```bash
   go run main.go
   ```

4. **Access the API at:** `http://localhost:8080`
