package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Println("New connection accepted")
	response := "HTTP/1.1 200 OK\r\n\r\n"
	fmt.Println("Attempting to send response:", response)
	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing to connection:", err.Error())
	} else {
		fmt.Println("Response sent successfully")
	}
}
