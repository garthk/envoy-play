package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "second\n")
}

func main() {
	http.HandleFunc("/second", handler)
	addr := getBindAddr()
	fmt.Fprintf(os.Stderr, "Listening on %s\n", addr)
	http.ListenAndServe(addr, nil)
}

func getBindAddr() string {
	addr := getEnv("SVC_ADDR", "127.0.0.1")
	port, _ := strconv.ParseInt(getEnv("SVC_PORT", "8082"), 10, 16)
	return fmt.Sprintf("%s:%d", addr, port)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
