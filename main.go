package main

import (
	"io/ioutil"
	"fmt"
	"net/http"

	"github.com/op/go-logging"

	"github.com/cevaris/balance/lb"
)

var LB_ADDRESS string = "localhost:8000"
var lob *lb.LoadBalancer
var log = logging.MustGetLogger("balance")

func main() {
	logging.SetLevel(logging.INFO, "balance")

	lob = lb.NewLoadBalancer()
	mid := &lb.LoadBalancerMiddleware{Handler:http.HandlerFunc(handler)}
	server := &http.Server{
		Addr: LB_ADDRESS,
		Handler: mid,
	}
	fmt.Println("Listening on", LB_ADDRESS)
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Info("%s %s", r.Method, r.RequestURI)
	h := lob.NextHost()
	forwardRequest(w, r, h)
}

func forwardRequest(w http.ResponseWriter, r *http.Request, h *lb.Host) {
	forwardURI := h.RequestURI(r.RequestURI)

	log.Debug("forwarding %s %s", r.Method, forwardURI)

	req, errReq := http.NewRequest(
		r.Method,
		forwardURI,
		r.Body)
	if errReq != nil {
		panic(errReq)
	}

	req.Header.Add("LB-RESPONDER", h.Addr)
	client := &http.Client{}
	resp, errResp := client.Do(req)
	if errResp != nil {
		panic(errResp)
	}

	defer resp.Body.Close()
	body, errBody := ioutil.ReadAll(resp.Body)
	if errBody != nil {
		panic(errBody)
	}
	w.Write(body)
}
