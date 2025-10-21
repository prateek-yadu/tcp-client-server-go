package main

import (
	"fmt"
	"net"
)

func main() {
	// create listener
	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		panic(err)
	}

	// will close the connection at the end
	defer listener.Close()

	fmt.Println("listening on 127.0.0.1:8000")

	// waits for connection forever
	for {
		// accepts the incomming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error Conneting", err)
			continue // skips this iteration and agin listen for connection
		}

		fmt.Println("Successfully Connected to: ", conn.RemoteAddr())

		go handleConnection(conn) 
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close() // closes connection at the end

	buffer := make([]byte, 1024) // creates buffer of 1kb

	// looks for message forever
	for {
		n, err := conn.Read(buffer) // reads data stores it in buffer and returns its lenght in var n
		if n == 0 {                // checks if client is online or not if not exit.
			fmt.Println("Client Disconnected")
			return
		}
		if err != nil {
			fmt.Println("Can not read the message", err)
			continue
		}
		fmt.Println("Recived from client", string(buffer[:n])) // parse buffer data in string
		conn.Write([]byte("Hello world"))
	}

}
