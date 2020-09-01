package main

import (
	"fmt"
	"net/http"
)

func df(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Mohit Jindal",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Toshi Jindal",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)

	fmt.Fprintln(w, "Cookie is set")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
	if c1, err := r.Cookie("first_cookie"); err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	} else {
		fmt.Fprintln(w, c1)
	}
	cs := r.Cookies()
	fmt.Fprintln(w, cs[0])
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/", df)
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
