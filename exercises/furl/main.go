// Minimal curl.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func fetch(link string, wg *sync.WaitGroup) {
	// (2) Fetch link, measure how long the
	// request took and write elapsed time to
	// stdout.
	defer wg.Done()
	started := time.Now() // elapsed := time.Since(started)
	resp, err := http.Get(link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Printf("%s\t%s\n", time.Since(started), link)
}

func main() {
	links := []string{
		"http://devopenspace.de",
		"https://l.de",
		"https://www.heise.de/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://ard.de/",
		"https://zdf.de/",
		"https://lvz.de/",
	}
	var wg sync.WaitGroup
	// (1) Fetch all links
	for _, link := range links {
		wg.Add(1)
		go fetch(link, &wg)
	}
	wg.Wait()
	// time.Sleep(2 * time.Second)
}
