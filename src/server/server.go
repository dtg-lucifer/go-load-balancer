package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ServerDetails struct {
	address string
	proxy   *httputil.ReverseProxy
}

func NewServer(address string) *ServerDetails {
	serverUrl, err := url.Parse(address)
	if err != nil {
		panic(err)
	}
	return &ServerDetails{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func (s *ServerDetails) Address() string {
	return s.address
}

func (s *ServerDetails) IsAlive() bool {
	return true
}

func (s *ServerDetails) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}
