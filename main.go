package main

import (
	"io/ioutil"
	"fmt"
	"net/http"

	"github.com/cevaris/balance/lb"
)

var LB_ADDRESS string = "localhost:8000"
var lob *balance.LoadBalancer


func main() {

	lob = balance.NewLoadBalancer()
	mid := &balance.LoadBalancerMiddleware{Handler:http.HandlerFunc(index)}
	server := &http.Server{
		Addr: LB_ADDRESS,
		// Handler: logs.NewApacheLoggingHandler(mux, os.Stdout),
		Handler: mid,
	}
	fmt.Println("Listening on", LB_ADDRESS)
	server.ListenAndServe()

}

func index(w http.ResponseWriter, r *http.Request) {

	 host := lob.RandHost()

	fmt.Println(r.Method)
	fmt.Println(r.RequestURI)

	req, errReq := http.NewRequest(
		r.Method,
		host.RequestURI(r.RequestURI),
		r.Body)
	if errReq != nil {
		panic(errReq)
	}

	req.Header.Add("LB-RESPONDER", host.Addr)
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
