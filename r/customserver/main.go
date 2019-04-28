// Customize server.
package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

func myHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func main() {
	s := &http.Server{
		Addr:           ":8080",
		Handler:        http.HandlerFunc(myHandler),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("listening on http://localhost:8080")
	log.Fatal(s.ListenAndServe())
}
