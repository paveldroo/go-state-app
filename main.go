package main

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

var cnt = 0

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/set", set)
	http.HandleFunc("/read", read)
	http.HandleFunc("/expire", expire)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func index(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("my-cookie")
	if errors.Is(err, http.ErrNoCookie) {
		io.WriteString(w, `<a href="/set">Set Cookie</>`)
	}
	io.WriteString(w, `<a href="/read">Read Cookie</>`)
}

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{Name: "my-cookie", Value: "SomeWeirdCookieHere"})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func read(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if errors.Is(err, http.ErrNoCookie) {
		http.Redirect(w, r, "/", 303)
	}
	io.WriteString(w, "Your cookie is: "+c.Value)
}

func expire(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if errors.Is(err, http.ErrNoCookie) {
		http.Redirect(w, r, "/", 303)
	}
	c.Expires = time.Now()
	http.SetCookie(w, c)
}
