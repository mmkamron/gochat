package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

var (
	clients      = make(map[net.Conn]string)
	messageQueue = make(chan string)
)

func main() {
	listener, err := net.Listen("tcp", ":"+os.Getenv("PORT"))
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	defer listener.Close()

	fmt.Println("Server is running. Accepting connections...")

	go broadcastMessages()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept client connection:", err)
			continue
		}

		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)

	// Read the username from the client
	username, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Failed to read username from client:", err)
		conn.Close()
		return
	}

	username = username[:len(username)-1] // Remove the newline character

	clients[conn] = username
	messageQueue <- fmt.Sprintf("User '%s' has joined the chat.", username)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				continue
			} else {
				log.Println("Failed to read message from client:", err)
				break
			}
		}

		messageQueue <- fmt.Sprintf("%s: %s", username, strings.TrimSpace(message))
	}

	delete(clients, conn)
	messageQueue <- fmt.Sprintf("User '%s' has left the chat.", username)
	conn.Close()
}

func broadcastMessages() {
	for message := range messageQueue {
		for client := range clients {
			_, err := client.Write([]byte(message + "\n"))
			if err != nil {
				log.Println("Failed to send message to client:", err)
			}
		}
	}
}
