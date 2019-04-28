package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/bye/", bye)
	http.HandleFunc("/", home)

	log.Println("listening on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// hello works with a query parameter, name.
func hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Hello %s", name)
}

// bye assumes a path like /bye/berlin.
func bye(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.URL.Path, "/")
	name := parts[2]
	if name == "" {
		name = "World"
	}
	fmt.Fprintf(w, "Bye %s", name)
}

// home handles the homepage.
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// Content type is inferred, takes a bit of time. Set explicitly.
	io.WriteString(w, `<html><body>Homepage:
		<a href="/hello?name=World">/hello?name=World</a>,
		<a href="/bye/berlin">/bye/berlin</a>`)
}
