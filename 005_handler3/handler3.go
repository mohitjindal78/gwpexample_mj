package main

import (
	"fmt"
	"net/http"
)

type DfHandler struct{}

func (h *DfHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

type WorldHandler struct{}

func (h *WorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	hellohandler := HelloHandler{}
	worldhandler := WorldHandler{}
	dfhandler := DfHandler{}

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.Handle("/", &dfhandler)
	http.Handle("/hello", &hellohandler)
	http.Handle("/world", &worldhandler)

	server.ListenAndServe()
}
