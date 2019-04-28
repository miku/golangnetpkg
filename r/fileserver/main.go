package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("server at http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/usr/share/doc"))))
}
