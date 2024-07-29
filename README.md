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


Made with the help of [CodeCrafters](https://app.codecrafters.io/courses/http-server/introduction).