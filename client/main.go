package main

import (
	"fmt"
	"log"
	"net"
)

const (
	serverProtocol = "tcp"
	serverAddress  = "127.0.0.1"
	serverPort     = ":8080"
)

func main() {
	connection, err := net.Dial(serverProtocol, serverAddress+serverPort)
	if err != nil {
		log.Fatal("cannot establish a connection to server", err)
	}
	defer connection.Close()

	err = sendMessage(connection, "Hello server socket, I'm the client socket!")
	if err != nil {
		log.Println(err)
		return
	}

	serverResponse, err := readResponse(connection)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Printf("Received from server: %s\n", serverResponse)
}

func sendMessage(connection net.Conn, message string) error {
	_, err := connection.Write([]byte(message))
	if err != nil {
		return fmt.Errorf("cannot write client data to connection: %w", err)
	}
	return nil
}

func readResponse(connection net.Conn) (string, error) {
	buffer := make([]byte, 1024)
	serverMsgLen, err := connection.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("cannot read server response: %w", err)
	}
	return string(buffer[:serverMsgLen]), nil
}
