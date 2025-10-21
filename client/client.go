package main

import (
	"fmt"
	"net"
)

func main() {
	// connects to the server
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
	defer conn.Close() // closes cnnection at the end

	// sends message to server


	for { // stays connected forever

		conn.Write([]byte("Hello from client")) // writes message for server

		// creates buffer byte of 1kb
		buffer := make([]byte, 1024)

		// reads data from server
		n, _ := conn.Read(buffer)
		if n == 0 { // checks if client is online or not if not exit.
			fmt.Println("Server Disconnected")
			break
		}

		fmt.Println("Server said", string(buffer[:n]))
	}
}
