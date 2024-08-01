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

Send an HTTP `GET` request to your server:

```bash
$ curl -v http://localhost:4221
```

Server responds to the request with the following response:

```javascript
HTTP/1.1 200 OK\r\n\r\n
```

### 3. Extract URL Path

Send a `GET` request, with a random string as the path:

```bash
$ curl -v http://localhost:4221/abcdefg
```

Server responds to this request with a `404` response:

```javascript
HTTP/1.1 404 Not Found\r\n\r\n
```

Send a `GET` request, with the path `/`:

```bash
$ curl -v http://localhost:4221
```

Server responds to this request with a `200` response:

```javascript
HTTP/1.1 200 OK\r\n\r\n
```

### 4. Respond with body

Send a `GET` request to the `/echo/{str}` endpoint on the server, with some random string.

```bash
$ curl -v http://localhost:4221/echo/abc
```

Server responds with a `200` response that contains the following parts:

- `Content-Type` header set to `text/plain`.
- `Content-Length` header set to the length of the given string.
- Response body set to the given string.

```javascript
HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 3\r\n\r\nabc
```
### 5. Read header

Send a `GET` request to the `/user-agent` endpoint on the server. The request will have a `User-Agent` header.

```bash
$ curl -v --header "User-Agent: foobar/1.2.3" http://localhost:4221/user-agent
```

Server responds with a `200` response that contains the following parts:

- `Content-Type` header set to `text/plain`.
- `Content-Length` header set to the length of the `User-Agent` value.
- Message body set to the `User-Agent` value.

```javascript
HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 12\r\n\r\nfoobar/1.2.3
```