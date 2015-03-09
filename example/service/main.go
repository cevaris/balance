package main

import (
	"io"
	"flag"
	"fmt"
	"math/rand"
	"net/http"
)


var host string
var port int

func location(h string, p int) string {
	return fmt.Sprintf("%s:%d", h, p)
}

func main() {
	flag.IntVar(&port, "p", 5000, "port number")
	flag.StringVar(&host, "h", "localhost", "host location")
	flag.Parse()

	mux := http.DefaultServeMux
	mux.HandleFunc("/", index)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/whoami", whoami)
	mux.HandleFunc("/rand-int", randInt)
	server := &http.Server{
		Addr: location(host, port),
	}

	fmt.Println("Listening on", location(host, port))
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome to golang!")
}

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func whoami(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("I am %s",location(host, port)))
}

func randInt(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("%d",rand.Int31()))
}
