package balance

import (
	"fmt"
	"math/rand"
	"net/http"
)

/*
http://www.peplink.com/technology/load-balancing-algorithms/
*/

type LoadBalancer struct {
	Hosts []*Host
}


type Host struct {
	Addr string
	Weight int
}

func (h *Host) RequestURI(path string) string {
	return fmt.Sprintf("http://%s%s", h.Addr, path)
}

func NewHost(addr string, weight int) *Host {
	return &Host{Addr: addr, Weight: weight}
}

func NewLoadBalancer() *LoadBalancer {
	h1 := NewHost("localhost:5000", 0)
	h2 := NewHost("localhost:5001", 1)
	h3 := NewHost("localhost:5002", 2)
	return &LoadBalancer{
		Hosts: []*Host{h1, h2, h3},
	}
}


func (l *LoadBalancer) NextHost() *Host {
	return l.RandHost()
}

func (l *LoadBalancer) RandHost() *Host {
	hostCount := len(l.Hosts)
	return l.Hosts[rand.Intn(hostCount)]
}


type LoadBalancerMiddleware struct {
	Handler http.Handler
}

func (a *LoadBalancerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.Handler.ServeHTTP(w, r)
	w.Write([]byte("<!-- Middleware says hello! -->"))
}



