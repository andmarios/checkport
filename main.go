package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

var (
	address = flag.String("address", "0.0.0.0", "IP address to try to bind to")
	port    = flag.String("port", "80", "port to try bind to")

//	protocol = flag.String("proto", "tcp", "protocol to use")
)

func main() {
	flag.Parse()

	conn, err := net.Listen("tcp", *address+":"+*port)
	if err != nil {
		fmt.Println("Error binding to port:", err.Error())
		os.Exit(1)
	}

	defer conn.Close()
}
