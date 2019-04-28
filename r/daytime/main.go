package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	hostport := "127.0.0.1:1200"
	log.Printf("listening on %s", hostport)
	tcpAddr, err := net.ResolveTCPAddr("tcp4", hostport)
	if err != nil {
		log.Fatal(err)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		daytime := time.Now().String()
		if _, err := io.WriteString(conn, daytime); err != nil {
			log.Fatal(err)
		}
		conn.Close()
	}
}
