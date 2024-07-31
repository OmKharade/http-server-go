package main

import (
	"fmt"
	"net"
	"os"
	"strings"
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

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err.Error())
		return
	}

	request := string(buffer[:n])
	requestLine := strings.Split(request, "\n")[0]
	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(parts) != 3 {
		fmt.Println("Invalid HTTP request")
		return
	}

	path := parts[1]
	fmt.Println("Requested path:", path)

	var response string
	if path == "/" {
		response = "HTTP/1.1 200 OK\r\n\r\n"
	} else if strings.HasPrefix(path, "/echo/") {
		echoStr := strings.TrimPrefix(path, "/echo/")
		contentLength := len(echoStr)
		response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type:text/plain\r\nContent-Length: %d\r\n\r\n%s", contentLength, echoStr)
	} else {
		response = "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	fmt.Println("Sending response:", response)
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing to connection:", err.Error())
	} else {
		fmt.Println("Response sent successfully")
	}
}
