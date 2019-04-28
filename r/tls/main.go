package main

import (
	// "fmt"
	// "io"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
	// fmt.Fprintf(w, "This is an example server.\n")
	// io.WriteString(w, "This is an example server.\n")
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServeTLS(":443", "example.com+5.pem", "example.com+5-key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
