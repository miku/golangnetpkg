package main

import (
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	l, err := net.Listen("tcp", "localhost:3030")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Would block.
		conn, err := l.Accept()
		if err != nil {
			// Maybe try again.
			log.Fatal(err)
		}
		// go copyToStderr(conn)
		go proxy(conn)
	}
}

func proxy(conn net.Conn) {
	defer conn.Close()

	remote, err := net.Dial("tcp", "spartakiade.de:80")
	if err != nil {
		log.Printf("dial failed: %v", err)
		return
	}
	defer remote.Close()

	// From browser to remote.
	go io.Copy(remote, conn) // OK to let run, as will exit on io.EOF anyway.

	// From remote to connection.
	io.Copy(conn, remote)
}

func copyToStderr(conn net.Conn) {
	defer conn.Close()
	// Custom version of copy, with a read timeout set.
	// n, err := io.Copy(os.Stderr, conn)
	for {
		conn.SetReadDeadline(time.Now().Add(5 * time.Second))
		var buf [128]byte
		n, err := conn.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("failed with err:  %v", err)
			return
		}
		os.Stderr.Write(buf[:n])
	}
}
