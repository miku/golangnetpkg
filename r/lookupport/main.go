// $ go run main.go
// service port: 9418
package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	port, err := net.LookupPort("tcp", "git")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("service port: %d\n", port)
}
