package challenges

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const rawCodeReviewSnippet = `package main

import (
	"fmt"
	"os"
	"time"

	curl "github.com/andelf/go-curl"
)

func writeCallback(buf []byte, userdata interface{}) bool {
	// Simply throw the data away.
	return true
}

func timedRetrieve(curlInit *curl.CURL, urlStr string) float64 {
	curlInit.Setopt(curl.OPT_URL, urlStr)
	curlInit.Setopt(curl.OPT_WRITEFUNCTION, writeCallback)

	// Do the request
	start := time.Now()
	_ = curlInit.Perform()
	duration := time.Since(start)
	// See what response code we got
	res, _ := curlInit.Getinfo(curl.INFO_RESPONSE_CODE)
	if resAsInt, ok := res.(int); ok {
		if resAsInt >= 400 {
			return -1
		}
		return duration.Seconds()
	}
	return -1
}

func retrieveAll(urls []string, sum []float64, count []int) ([]string, []float64, []int) {
	curlInit := curl.EasyInit()
	defer curlInit.Cleanup()

	for i := 0; i < len(urls); i++ {
		latency := timedRetrieve(curlInit, urls[i])
		if latency < 0 {
			// Retrieval failed, stop trying to poll this URL.
			urls = append(urls[:i], urls[i+1:]...)
			i--
		} else {
			sum[i] += latency
			count[i]++
		}
	}

	return urls, sum, count
}

func main() {
	// Read the list of URLs from command line arguments
	urls := os.Args[1:]
	sum, count := []float64{}, []int{}
	for range urls {
		sum = append(sum, 0)
		count = append(count, 0)
	}

	for i := 0; i < 10; i++ {
		urls, sum, count = retrieveAll(urls, sum, count)
		time.Sleep(time.Second)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(urls[i], ":", (sum[i] / float64(count[i])), "ms")
	}
}`

func CodeReviewEndpoint(c *gin.Context) {
	issues := []string{
		"🔴 Modifica slices mientras los recorres en `retrieveAll`: esto puede causar errores inesperados.",
		"🟠 No se manejan errores del todo correctamente (e.g., `Perform()` y `Getinfo()` los ignora).",
		"🟡 El uso compartido de `curl.EasyInit()` entre múltiples URLs puede ser riesgoso si fuera concurrente.",
		"🔵 Se podría usar `make([]float64, len(urls))` en lugar de `append` en el bucle.",
		"🟢 El código no es modular: se puede separar en paquetes o funciones reutilizables.",
		"🟣 Usa `os.Args` en lugar de entrada desde stdin o HTTP, lo que lo hace menos flexible.",
		"⚫ No hay logs ni trazas para depurar errores.",
		"🟤 `return -1` como señal de error no es expresivo: mejor retornar `(float64, error)`.",
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    rawCodeReviewSnippet,
		"review":  strings.Join(issues, "\n"),
		"summary": "Se identificaron varios problemas estructurales, de legibilidad y robustez. Se recomienda modularizar, mejorar manejo de errores y evitar modificar slices mientras se recorren.",
	})
}
