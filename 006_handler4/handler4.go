package main

import (
	"fmt"
	"net/http"
)

func df(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World!")
	h := r.Header
	//fmt.Fprintln(w, "\n")
	fmt.Fprintln(w, h)
	fmt.Fprintln(w, h["Accept"])
	fmt.Fprintln(w, h["Accept"][0])
	fmt.Fprintln(w, h.Get("Accept"))
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}

	http.HandleFunc("/", df)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)

	server.ListenAndServe()
}
