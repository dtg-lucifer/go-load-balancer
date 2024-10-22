package loadbalancer

import (
	"fmt"
	"net/http"

	"go-load-balancer/src/server"
)

type LoadBalancer struct {
	servers         []server.Server
	Port            string
	roundRobinCount int
}

func NewLoadBalancer(port string, servers []server.Server) *LoadBalancer {
	return &LoadBalancer{
		Port:            port,
		servers:         servers,
		roundRobinCount: 0,
	}
}

func (lb *LoadBalancer) GetNextAvailableServer() server.Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) ServeProxy(rw http.ResponseWriter, r *http.Request) {
	target := lb.GetNextAvailableServer()
	if target == nil {
		http.Error(rw, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	fmt.Printf("Redirecting request to: %s\n", target.Address())
	target.Serve(rw, r)
}
