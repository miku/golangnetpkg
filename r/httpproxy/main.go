// Minimal proxy example. HTTP only site: http://neuschwanstein.de/
package main

import (
	"flag"
	"io"
	"log"
	"net/http"
)

var addr = flag.String("addr", "127.0.0.1:8080", "hostport to listen on")

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

type proxy struct{}

func (p *proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, " ", r.Method, " ", r.URL)

	if r.URL.Scheme != "http" {
		msg := "unsupported protocal scheme " + r.URL.Scheme
		http.Error(w, msg, http.StatusBadRequest)
		log.Println(msg)
		return
	}

	client := &http.Client{}

	//http: Request.RequestURI can't be set in client requests.
	//http://golang.org/src/pkg/net/http/client.go
	log.Println(r.RequestURI)
	r.RequestURI = ""

	resp, err := client.Do(r)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Fatal("ServeHTTP:", err)
	}
	defer resp.Body.Close()

	log.Println(r.RemoteAddr, " ", resp.Status)

	copyHeader(w.Header(), resp.Header)
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func main() {
	flag.Parse()

	handler := &proxy{}

	log.Println("Starting proxy server on", *addr)
	log.Fatal(http.ListenAndServe(*addr, handler))
}
