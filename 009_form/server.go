package main

import (
	"fmt"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//Reading header
	//fmt.Fprintln(w, "")
	//fmt.Fprintln(w, r.Header)

	//Reading Body
	//len := r.ContentLength
	//body := make([]byte, len)
	//r.Body.Read(body)
	//fmt.Fprintln(w, string(body))

	//Reading form type
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
	fmt.Fprintln(w, r.PostForm)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
