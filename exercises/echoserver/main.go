package main

import (
	"log"
	"net"
)

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		// (3) Read from conn.
		// (4) Write to conn.
	}
}

func main() {
	hostport := "localhost:1201"

	// (1) Listen in TCP.
	for {
		// (2) Accept connection
	}
}
