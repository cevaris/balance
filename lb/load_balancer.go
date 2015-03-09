package lb

import (
	"fmt"
	"math/rand"
	"net/http"
)

// http://www.peplink.com/technology/load-balancing-algorithms/
const (
	RandomStrategy = iota
	WeightedStrategy
	PriorityStrategy
)


type Host struct {
	Addr string
	Weight int
	WeightPercent int
}

func (h *Host) RequestURI(path string) string {
	return fmt.Sprintf("http://%s%s", h.Addr, path)
}

func NewHost(addr string, weight int) *Host {
	return &Host{
		Addr: addr,
		Weight: weight,
		WeightPercent: 1.0,
	}
}



type LoadBalancer struct {
	Hosts []*Host
}

func NewLoadBalancer() *LoadBalancer {
	h1 := NewHost("localhost:5000", 1)
	h2 := NewHost("localhost:5001", 2)
	h3 := NewHost("localhost:5002", 3)
	hosts := []*Host{h1, h2, h3}

	lob := &LoadBalancer{}
	lob.Hosts = setupWeightPercentages(hosts)

	return lob
}

func setupWeightPercentages(hosts []*Host) []*Host{
	weights := make([]int, len(hosts))
	for i, v := range(hosts) {
		weights[i] = v.Weight
	}
	percents := calcWeightPercentages(weights)
	for i, v := range(hosts) {
		v.WeightPercent = percents[i]
	}
	return hosts
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



