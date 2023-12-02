package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	fmt.Println("Initial setup")

	l, err := net.Listen("tcp", "0.0.0.0:6379")

	if err != nil {
		fmt.Println("Failed to bind to port 6379", err.Error())
		os.Exit(1)
	}

	conn, err := l.Accept()

	if err != nil {
		fmt.Println("Error Accepting connection", err.Error())
		os.Exit(1)
	}

	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
}
