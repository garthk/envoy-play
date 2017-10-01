package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Host configures a host to which we will direct traffic.
type Host struct {
	URL string `json:"url"`
}

// Hosts configures hosts which we will direct traffic.
type Hosts []Host

// Cluster configures a cluster of hosts to which we will direct traffic.
// It is NOT comprehensive vs Envoy's configurtion schemata.
type Cluster struct {
	Name             string `json:"name"`
	Type             string `json:"type"`
	ConnectTimeout   int    `json:"connect_timeout_ms"`
	LoadBalancerType string `json:"lb_type"`
	Hosts            `json:"hosts"`
}

// CDSResponse configures clusters of hosts to which we will direct traffic.
type CDSResponse struct {
	Clusters []Cluster `json:"clusters"`
}

func clusters(c echo.Context) error {
	serviceCluster := c.Param("service_cluster")
	serviceNode := c.Param("service_node")
	fmt.Fprintf(os.Stderr, "clusters for service cluster %s node %s\n", serviceCluster, serviceNode)
	return c.JSON(http.StatusOK, CDSResponse{
		Clusters: []Cluster{
			Cluster{
				Name:             "first",
				Type:             "strict_dns",
				ConnectTimeout:   100,
				LoadBalancerType: "least_request",
				Hosts: Hosts{Host{
					URL: "tcp://first:8081",
				}},
			}, Cluster{
				Name:             "second",
				Type:             "strict_dns",
				ConnectTimeout:   100,
				LoadBalancerType: "least_request",
				Hosts: Hosts{Host{
					URL: "tcp://second:8082",
				}},
			},
		},
	})
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.GET("/v1/clusters/:service_cluster/:service_node", clusters)

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
