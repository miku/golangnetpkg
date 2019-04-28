package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	hostport := "127.0.0.1:1200"
	log.Println(hostport)

	listener, err := net.Listen("tcp", hostport)
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
