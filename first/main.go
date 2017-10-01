package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

func mkhandler(targetURL string) func(http.ResponseWriter, *http.Request) {
	var netClient = &http.Client{
		Timeout: time.Second * 2,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "first\n")
		response, err := netClient.Get(targetURL)
		if err != nil {
			panic(err)
		}
		io.Copy(w, response.Body)
		response.Body.Close()
	}
}

func main() {
	targetURL := getEnv("SVC_TARGET_URL", "http://127.0.0.1:8082/second")
	fmt.Fprintf(os.Stderr, "Target is: %s\n", targetURL)
	http.HandleFunc("/first", mkhandler(targetURL))
	addr := getBindAddr()
	fmt.Fprintf(os.Stderr, "Listening on %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func getBindAddr() string {
	addr := getEnv("SVC_ADDR", "127.0.0.1")
	port, _ := strconv.ParseInt(getEnv("SVC_PORT", "8081"), 10, 16)
	return fmt.Sprintf("%s:%d", addr, port)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
