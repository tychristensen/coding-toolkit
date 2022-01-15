package tcpserver

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// StartServer starts a simple server process that listens for TCP connections
// on port localhost:9500 and returns a funny resonse on the connection
func StartServer(readyChan chan bool) {
	// <-readyChan
	println("server started")
	listener, err := net.Listen("tcp", ":9500")
	if err != nil {
		panic(fmt.Errorf("failed to make listener: %v", err))
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("failed to accept connection: %v", err)
			continue
		}

		log.Println("Accepted", conn.RemoteAddr())
		conn.Write([]byte(">"))
		go handleConn(conn)
	}
}

// handleConn handles the new connection accepted in StartServer
func handleConn(conn net.Conn) {
	defer conn.Close()

	s := bufio.NewScanner(conn)

	for s.Scan() {
		data := s.Text()

		if data == "" {
			conn.Write([]byte(">"))
			continue
		}

		if data == "exit" {
			return
		}

		conn.Write([]byte("damn boi\n>"))
	}
}
