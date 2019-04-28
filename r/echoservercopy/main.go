package main

import (
	"io"
	"log"
	"net"
)

func handleClient(conn net.Conn) {
	if _, err := io.Copy(conn, conn); err != nil {
		log.Fatal(err)
	}
}

func main() {
	hostport := "localhost:1201"
	log.Printf("listening on %s", hostport)
	listener, err := net.Listen("tcp", hostport)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		handleClient(conn)
		conn.Close()
	}
}
