package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net"
	"os"
	"path/filepath"
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
	lines := strings.Split(request, "\r\n")
	if len(lines) < 1 {
		fmt.Println("Invalid HTTP request")
		return
	}

	requestLine := lines[0]
	parts := strings.Split(strings.TrimSpace(requestLine), " ")
	if len(parts) != 3 {
		fmt.Println("Invalid HTTP request")
		return
	}
	method := parts[0]
	path := parts[1]
	fmt.Println("Requested method:", method)
	fmt.Println("Requested path:", path)

	userAgent := ""
	acceptEncoding := ""
	for _, line := range lines[1:] {
		if strings.HasPrefix(line, "User-Agent: ") {
			userAgent = strings.TrimPrefix(line, "User-Agent: ")
		} else if strings.HasPrefix(line, "Accept-Encoding: ") {
			encoding := strings.Split(strings.TrimPrefix(line, "Accept-Encoding: "), ",")
			for _, scheme := range encoding {
				if strings.TrimSpace(scheme) == "gzip" {
					acceptEncoding = "gzip"
				}
			}
		}
	}
	var response string
	if path == "/" {
		response = "HTTP/1.1 200 OK\r\n\r\n"
	} else if strings.HasPrefix(path, "/echo/") {
		echoStr := strings.TrimPrefix(path, "/echo/")
		if acceptEncoding == "gzip" {
			var buf bytes.Buffer
			gzipWriter := gzip.NewWriter(&buf)
			_, err := gzipWriter.Write([]byte(echoStr))
			if err != nil {
				fmt.Println("Error compressing data:", err)
				response = "HTTP/1.1 500 Internal Server Error\r\n\r\n"
			} else {
				err = gzipWriter.Close()
				if err != nil {
					fmt.Println("Error closing gzip writer:", err)
					response = "HTTP/1.1 500 Internal Server Error\r\n\r\n"
				} else {
					compressedData := buf.Bytes()
					headers := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Encoding: gzip\r\nContent-Length: %d\r\n\r\n", len(compressedData))
					response = headers + string(compressedData)
				}
			}
		} else {
			contentLength := len(echoStr)
			response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", contentLength, echoStr)
		}
	} else if strings.HasPrefix(path, "/user-agent") {
		contentLength := len(userAgent)
		response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", contentLength, userAgent)
	} else if strings.HasPrefix(path, "/files/") {
		filename := strings.TrimPrefix(path, "/files/")
		dirname := os.Args[2]
		fullpath := filepath.Join(dirname, filename)
		if method == "GET" {
			data, err := os.ReadFile(fullpath)
			if err != nil {
				response = "HTTP/1.1 404 Not Found\r\n\r\n"
			} else {
				response = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n%s", len(data), data)
			}
		} else if method == "POST" {
			bodyStart := strings.Index(request, "\r\n\r\n")
			if bodyStart == -1 {
				response = "HTTP/1.1 400 Bad Request\r\n\r\n"
			} else {
				body := request[bodyStart+4:]
				err := os.WriteFile(fullpath, []byte(body), 0644)
				if err != nil {
					response = "HTTP/1.1 500 Internal Server Error\r\n\r\n"
				} else {
					response = "HTTP/1.1 201 Created\r\n\r\n"
				}
			}
		}
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
