package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Tags configures tags for a host
type Tags struct {
	AZ                  string `json:"az"`
	Canary              bool   `json:"canary"`
	LoadBalancingWeight int    `json:"load_balancing_weight"`
}

// Host configures a host to which we will direct traffic.
type Host struct {
	IPAddress string `json:"ip_address"`
	Port      int    `json:"port"`
	Tags      `json:"tags"`
}

// SDSResponse specifies hosts for a service
type SDSResponse struct {
	Hosts []Host `json:"hosts"`
}

// Cluster configures a cluster of hosts to which we will direct traffic.
// It is NOT comprehensive vs Envoy's configurtion schemata.
type Cluster struct {
	Name             string `json:"name"`
	ServiceName      string `json:"service_name"`
	Type             string `json:"type"`
	ConnectTimeout   int    `json:"connect_timeout_ms"`
	LoadBalancerType string `json:"lb_type"`
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
				ServiceName:      "first",
				Type:             "sds",
				ConnectTimeout:   100,
				LoadBalancerType: "least_request",
			}, Cluster{
				Name:             "second",
				ServiceName:      "second",
				Type:             "sds",
				ConnectTimeout:   100,
				LoadBalancerType: "least_request",
			},
		},
	})
}

func registration(c echo.Context) error {
	serviceName := c.Param("service_name")
	fmt.Fprintf(os.Stderr, "registrations for service name %s\n", serviceName)
	addrs, err := net.LookupHost(serviceName)
	if err != nil {
		return err
	}
	port, err := getPort(serviceName)
	if err != nil {
		return err
	}
	tags := Tags{
		AZ:                  "demo",
		Canary:              false,
		LoadBalancingWeight: 1,
	}
	hosts := make([]Host, len(addrs))
	for i, addr := range addrs {
		hosts[i] = Host{
			IPAddress: addr,
			Port:      port,
			Tags:      tags,
		}
	}
	return c.JSON(http.StatusOK, SDSResponse{
		Hosts: hosts,
	})
}

func getPort(serviceName string) (int, error) {
	switch serviceName {
	case "first":
		return 8081, nil
	case "second":
		return 8082, nil
	default:
		return 0, fmt.Errorf("getPort: unexpected serviceName %s", serviceName)
	}
}

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
