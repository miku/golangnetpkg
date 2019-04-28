package main

import (
	"fmt"
	"net/http"
)

type app struct{}

func (app app) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
}

func main() {
	http.Handle("/", app{})
	http.ListenAndServe(":8080", nil)
}
