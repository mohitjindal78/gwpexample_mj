package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintf(rw, "Hello world, %s!", rq.URL.Path[1:])
}
