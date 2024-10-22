package main

import (
	"fmt"
	"net/http"

	"go-load-balancer/src/loadbalancer"
	"go-load-balancer/src/server"
	"go-load-balancer/src/utils"
)

func main() {
	servers := []server.Server{
		server.NewServer("https://www.google.com"),
		server.NewServer("https://www.facebook.com"),
		server.NewServer("https://www.duckduckgo.com"),
	}

	lb := loadbalancer.NewLoadBalancer("8080", servers)

	http.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		lb.ServeProxy(rw, req)
	})

	fmt.Printf("Load Balancer listening on localhost:%s\n", lb.Port)
	utils.HandleError(http.ListenAndServe(":"+lb.Port, nil))
}
