package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type hello struct {
	Name string
	Date time.Time
}

func main() {
	payload := hello{
		Name: "World",
		Date: time.Now(),
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		enc := json.NewEncoder(w)
		if err := enc.Encode(payload); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	log.Printf("listening of http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
