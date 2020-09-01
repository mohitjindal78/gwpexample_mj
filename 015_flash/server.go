package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"
)

func df(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(302)
}

func setCookie(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Mohit Jindal!")
	c1 := http.Cookie{
		Name:     "flash",
		Value:    base64.URLEncoding.EncodeToString(msg),
		HttpOnly: true,
	}

	http.SetCookie(w, &c1)

	fmt.Fprintln(w, "Cookie is set")
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	if c, err := r.Cookie("flash"); err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "No cookie found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
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
