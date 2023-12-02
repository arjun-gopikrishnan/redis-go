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
		lines := strings.Split(strings.ReplaceAll(command, "\r", ""), `\n`)
		fmt.Print(lines)
		for _, line := range lines {
			if strings.ToUpper(line) == "PING" {
				conn.Write([]byte("+PONG\r\n"))
			} else {
				conn.Write([]byte("Command not recognized\r\n"))
			}
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
	}

}
