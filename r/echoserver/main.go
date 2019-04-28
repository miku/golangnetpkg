package main

import (
	"log"
	"net"
)

func handleClient(conn net.Conn) {
	var buf [512]byte
	for {
		n, err := conn.Read(buf[:])
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(buf[0:]))
		if _, err = conn.Write(buf[0:n]); err != nil {
			log.Println(err)
			return
		}
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
