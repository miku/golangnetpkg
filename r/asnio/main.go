package main

import (
	"encoding/asn1"
	"fmt"
	"log"
)

func main() {
	b, err := asn1.Marshal(42)
	if err != nil {
		log.Fatal(err)
	}
	var v int
	if _, err := asn1.Unmarshal(b, &v); err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)
}
