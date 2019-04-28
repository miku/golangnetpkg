package main

import (
	"fmt"
	"net/http"
)

func f(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
}

func main() {

	// Type adapter.
	handler := http.HandlerFunc(f)

	http.Handle("/", handler) // https://golang.org/pkg/net/http/#HandlerFunc
	http.ListenAndServe(":8080", nil)
}
