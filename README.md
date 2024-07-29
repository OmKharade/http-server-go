# http-server-go
Build your own HTTP Server using Go. 

### 1. Binding a port

Set up a basic TCP server that listens on port 4221.

It binds to the port and waits for a single incoming connection
```go
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
    _, err = l.Accept()
    if err != nil {
        fmt.Println("Error accepting connection: ", err.Error())
        os.Exit(1)
    }
}
```

### 2. Respond with 200

Accept concurrent connections
```go
for {
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection:", err.Error())
		continue
	}
	go handleConnection(conn)
}
```

Write data (response) to connection
```go
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
```
