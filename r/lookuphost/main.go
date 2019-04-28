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
	// Returns a list of strings.
	addrs, err := net.LookupHost(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	for _, addr := range addrs {
		fmt.Printf("%s\n", addr)
	}
}
