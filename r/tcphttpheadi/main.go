// $ go run main.go heise.de:80
// HTTP/1.1 301 Moved Permanently
// Date: Thu, 25 Apr 2019 12:41:19 GMT
// Server: Apache
// X-Cobbler: servo65.heise.de
// X-Pect: The Spanish Inquisition
// X-Clacks-Overhead: GNU Terry Pratchett
// X-42: DON'T PANIC
// Location: https://www.heise.de/
// Connection: close
// Content-Type: text/html; charset=iso-8859-1
//
package main

import (
	"flag"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		log.Fatal("host:port required")
	}

	conn, err := net.Dial("tcp", flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// HTTP/1.1 would require the Host header to be set,
	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Host.
	if _, err := io.WriteString(conn, "HEAD / HTTP/1.0\r\n\r\n"); err != nil {
		log.Fatal(err)
	}
	//b, err := ioutil.ReadAll(conn)
	//if err != nil {
	//	log.Fatal(err)
	//}
	if _, err := io.Copy(os.Stdout, conn); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(b))
}
