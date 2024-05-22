package main

import (
	"fmt"
	"log"
	"net"
)

const (
	protocol = "tcp"
	address  = "127.0.0.1"
	port     = ":8080"
)

func main() {
	listener, err := net.Listen(protocol, address+port)
	if err != nil {
		log.Fatal("cannot open server socket", err)
	}
	defer listener.Close()
	log.Printf("Server socket is listening on %s%s\n", address, port)

	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("cannot accept client connection", err)
			continue
		}

		go func(connection net.Conn) {
			err = handleConnection(connection)
			if err != nil {
				log.Println(err)
			}
		}(connection)
	}
}

func handleConnection(connection net.Conn) error {
	buffer := make([]byte, 1024)
	clientMsgLen, err := connection.Read(buffer)
	if err != nil {
		return err
	}
	log.Printf("Received from client: %s\n", string(buffer[:clientMsgLen]))

	err = sendResponse(connection, "Server socket says hello!")
	if err != nil {
		return err
	}
	return connection.Close()
}

func sendResponse(connection net.Conn, message string) error {
	_, err := connection.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("cannot write server data to connection: %w", err)
	}
	return nil
}
