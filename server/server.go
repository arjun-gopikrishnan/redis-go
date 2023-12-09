package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type Server struct {
	listener net.Listener
	port     string
}

func NewServer(address string) (*Server, error) {
	listener, err := net.Listen("tcp", address)

	if err != nil {
		return nil, err
	}

	return &Server{listener: listener, port: address}, nil
}

func (s *Server) Start() {
	fmt.Printf("Server is started and listening on port %s", s.port)

	for {
		conn, err := s.listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection", err.Error())
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

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

func (s *Server) Close() {
	s.listener.Close()
	fmt.Println("Server is stopped")
}
