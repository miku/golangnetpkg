package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatalf("usage: %s HOSTNAME", os.Args[0])
	}
	addr, err := net.ResolveIPAddr("ip", flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(addr)
}
