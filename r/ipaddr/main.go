// Validate IP address on the command line.
// $ go run main.go 0::1
// address is [00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000001] ::1
//
// $ go run main.go 1.1.1.1
// address is [00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 00000000 11111111 11111111 00000001 00000001 00000001 00000001] 1.1.1.1
//
// You see the v4InV6Prefix, https://git.io/fjs8I
package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func main() {
	flag.Parse()

	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: %s IP\n", os.Args[0])
		os.Exit(1)
	}
	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Printf("invalid address: %s\n", name)
	} else {
		fmt.Printf("address is %08b %s\n", addr, addr)
	}
}
