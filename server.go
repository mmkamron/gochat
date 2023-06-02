package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	message string
	addr    string
}

var (
	clients      = make(map[net.Conn]bool)
	messageQueue = make(chan client)
)

func main() {
	listener, err := net.Listen("tcp", ":1337")
	if err != nil {
		log.Fatal("Failed to start server:", err)
	}
	defer listener.Close()

	go broadcastMessages()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept client connection:", err)
			continue
		}

		clients[conn] = true
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to read message from client:", err)
			break
		}
		messageQueue <- client{conn.RemoteAddr().String(), message}
	}

	delete(clients, conn)
	conn.Close()
}

func broadcastMessages() {
	for message := range messageQueue {
		for client := range clients {
			_, err := client.Write([]byte(fmt.Sprintf("%s: %s", message.addr, message.message)))
			if err != nil {
				log.Println("Failed to send message to client:", err)
			}
		}
	}
}
