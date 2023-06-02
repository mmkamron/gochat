package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":1337")
	if err != nil {
		log.Fatal("Failed to connect to server:", err)
	}
	defer conn.Close()

	go receiveMessages(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")

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

		fmt.Println(message)
	}
}
