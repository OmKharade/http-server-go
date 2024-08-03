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

### 6. Return a file

Execute the program with a `--directory` flag. The `--directory` flag specifies the directory where the files are stored, as an absolute path.

```shell
$ ./your_program.sh --directory /tmp/
```

Send two `GET` requests to the `/files/{filename}` endpoint on your server.

#### First request

Ask for a file that exists in the files directory:

```shell
$ echo -n 'Hello, World!' > /tmp/foo
$ curl -i http://localhost:4221/files/foo
```

Server responds with a `200` response that contains the following parts:

- `Content-Type` header set to `application/octet-stream`.
- `Content-Length` header set to the size of the file, in bytes.
- Response body set to the file contents.

```js
HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: 14\r\n\r\nHello, World!
```

#### Second request

Ask for a file that doesn't exist in the files directory:

```shell
$ curl -i http://localhost:4221/files/non_existant_file
```

Server responds with a `404` response:

```js
HTTP/1.1 404 Not Found\r\n\r\n
```

### 7. Read request body

Execute your program with a `--directory` flag. The `--directory` flag specifies the directory to create the file in, as an absolute path.

```shell
$ ./your_program.sh --directory /tmp/
```

Send a `POST` request to the `/files/{filename}` endpoint on the server, with the following parts:

- `Content-Type` header set to `application/octet-stream`.
- `Content-Length` header set to the size of the request body, in bytes.
- Request body set to some random text.

```shell
$ curl -v --data "12345" -H "Content-Type: application/octet-stream" http://localhost:4221/files/file_123
```

Server returns a `201` response:

```js
HTTP/1.1 201 Created\r\n\r\n
```

Server also creates a new file in the files directory, with the following requirements:

- The filename must equal the `filename` parameter in the endpoint.
- The file must contain the contents of the request body.