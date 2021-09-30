package main

import "fmt"
import "sample/yamux"
import "net"
//import "os"

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

func main() {
    fmt.Println("Hello, World!")    
    client()
}

func client() {
    // Get a TCP connection
    conn, err := net.Dial(CONN_TYPE,CONN_HOST+":"+CONN_PORT)
    if err != nil {
        panic(err)
    }

    // Setup client side of yamux
    session, err := yamux.Client(conn, nil)
    if err != nil {
        panic(err)
    }

    // Open a new stream
    stream, err := session.Open()
    if err != nil {
        panic(err)
    }

    // Stream implements net.Conn
    stream.Write([]byte("ping"))
}
