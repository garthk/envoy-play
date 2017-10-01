package main

import (
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/labstack/echo"
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

// Respond to requests for service host details
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
