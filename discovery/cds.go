package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

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

// Respond to requests for cluster details
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
