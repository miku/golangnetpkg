package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Index")
	})

	r.HandleFunc("/a/{title}/p/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
