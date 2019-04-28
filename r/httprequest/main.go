package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.Header.Write(os.Stderr)
	fmt.Fprintf(w, "Hello World, %s!", r.URL.Path[1:])
}
func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
