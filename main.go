package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		command := scanner.Text()

		if strings.ToUpper(command) == "PING" {
			conn.Write([]byte("+PONG\r\n"))
		} else {
			conn.Write([]byte("Command not recognized\r\n"))
		}

	}
}

func main() {
	fmt.Println("Initial setup")

	l, err := net.Listen("tcp", "0.0.0.0:6379")

	if err != nil {
		fmt.Println("Failed to bind to port 6379", err.Error())
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
		// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	}

	// conn, err := l.Accept()

	// if err != nil {
	// 	fmt.Println("Error Accepting connection", err.Error())
	// 	os.Exit(1)
	// }

	// fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
}
