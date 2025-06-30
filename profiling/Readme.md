# Profiling in Go with `pprof`

This guide explains how to use Go’s built-in profiler (`net/http/pprof`) to inspect and analyze your application behavior during runtime.

---

## 📊 What can you profile?
- CPU usage
- Heap memory allocations
- Number of goroutines
- Blocking operations
- Mutex contention

---

## ⚙️ How to enable `pprof`
Add the following to your `main.go`:

```go
import _ "net/http/pprof"
go http.ListenAndServe(":6060", nil)
```
> ⚠️ Only use in development or behind secure access in production.

---

## 🚀 How to profile

1. **Run the app with pprof enabled**:
```bash
make run
```

2. **Capture a CPU profile (30 seconds)**:
```bash
make profile
```

3. **Check active goroutines**:
```bash
make goroutines
```

4. **Analyze the captured profile**:
```bash
go tool pprof cpu.prof
(pprof) top
(pprof) web    # requires graphviz
```

---

## 📁 Related files
- `Makefile` – contains profiling automation commands
- `pprof` – native Go profiling endpoints

---

## 🔧 Optional Tools
- [Pyroscope](https://pyroscope.io/)
- [Grafana Cloud Profiles](https://grafana.com/oss/profiles/)
- [Google Cloud Profiler](https://cloud.google.com/profiler)

---