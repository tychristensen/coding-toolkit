package tcpserver

import (
	"fmt"
	"log"
	"net"
)

func StartServer(readyChan chan bool, msgChan chan string, address string) {
	<-readyChan
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
		conn.Write([]byte("damn boi\n"))
		conn.Close()
	}
}
