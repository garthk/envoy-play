package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/v1/clusters/:service_cluster/:service_node", clusters)
	e.GET("/v1/registration/:service_name", registration)

	addr := getBindAddr()
	fmt.Fprintf(os.Stderr, "Listening on %s\n", addr)
	e.Logger.Fatal(e.Start(addr))
}

func getBindAddr() string {
	addr := getEnv("SVC_ADDR", "127.0.0.1")
	port, _ := strconv.ParseInt(getEnv("SVC_PORT", "8083"), 10, 16)
	return fmt.Sprintf("%s:%d", addr, port)
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
