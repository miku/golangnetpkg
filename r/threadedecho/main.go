package main

import (
	"log"
	"net"
)

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte

	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			log.Println(err)
			return
		}
		if _, err := conn.Write(buf[0:n]); err != nil {
			log.Println(err)
			return
		}
	}
}

func main() {
	hostport := "localhost:1200"
	log.Printf("listening on %s", hostport)

	listener, err := net.Listen("tcp", hostport)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}
