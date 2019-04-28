package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	counter int
}

func (app *App) Index(w http.ResponseWriter, r *http.Request) {
	app.counter++
	fmt.Fprintf(w, "Index (%d)", app.counter)
}

func main() {
	app := &App{}
	r := mux.NewRouter()

	r.HandleFunc("/", app.Index)

	r.HandleFunc("/a/{title}/p/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]
		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	log.Fatal(http.ListenAndServe(":8080", r))
}
