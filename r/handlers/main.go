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
	app := app{}
	http.Handle("/", app) // https://golang.org/pkg/net/http/#Handler
	http.ListenAndServe(":8080", nil)
}
