package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

// Ensures gofmt doesn't remove the "net" and "os" imports in stage 1 (feel free to remove this!)
var _ = net.Listen
var _ = os.Exit

func main() {
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleMessagesForConnection(conn)
	}
}

func handleMessagesForConnection(connection net.Conn) {
	for {
		message := make([]byte, 256)
		_, err := connection.Read(message)
		if err != nil {
			fmt.Printf("Error reading message from connection %s\n", err.Error())
		}
		fmt.Println(message)

		_, err = connection.Write([]byte("+PONG\r\n"))

		if err != nil {
			fmt.Printf("Error writing message to connection %s\n", err.Error())
		}
	}

}
