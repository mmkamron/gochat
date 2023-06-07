package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1337")
	if err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
	defer conn.Close()

	// Register a username
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter username: ")
	username, _ := reader.ReadString('\n')
	username = strings.TrimSpace(username)

	// Send the username to the server
	_, err = conn.Write([]byte(username + "\n"))
	if err != nil {
		log.Fatal("Failed to send username to server:", err)
	}

	go receiveMessages(conn)

	fmt.Println("You are registered as:", username)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to read input:", err)
			continue
		}

		_, err = conn.Write([]byte(input))
		if err != nil {
			log.Println("Failed to send data to server:", err)
			continue
		}
	}
}

func receiveMessages(conn net.Conn) {
	reader := bufio.NewReader(conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to read message from server:", err)
			break
		}

		fmt.Print(message)
	}
}
