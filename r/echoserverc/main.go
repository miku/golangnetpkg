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
		go handleClient(conn)
	}
}
