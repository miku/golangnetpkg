// Minimal curl.
package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://devopenspace.de")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	log.Printf("StatusCode=%v\n", resp.StatusCode)
	log.Printf("Proto=%v", resp.Proto)
	resp.Header.Write(os.Stderr)

	if _, err := io.Copy(os.Stdout, resp.Body); err != nil {
		log.Fatal(err)
	}
}
