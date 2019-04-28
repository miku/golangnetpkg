package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	flag.Parse()

	var s = "127.0.0.1"
	if flag.NArg() > 0 {
		s = flag.Arg(0)
	}

	addr := net.ParseIP(s)
	if addr == nil {
		log.Fatal("invalid address")
	}

	mask := addr.DefaultMask() // Parse with e.g. net.CIDRMask(31, 32) or net.IPv4Mask(255, 255, 255, 0)
	network := addr.Mask(mask)
	ones, bits := mask.Size()

	fmt.Printf("addr:                   %s\n", addr)
	fmt.Printf("default mask length:    %d\n", bits)
	fmt.Printf("number of leading ones: %d\n", ones)
	fmt.Printf("mask (hex):             %s\n", mask)
	fmt.Printf("network:                %s\n", network)
}
