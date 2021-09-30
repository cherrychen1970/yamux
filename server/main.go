package main

import "fmt"
import "sample/yamux"
import "net"
import "os"

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

var quit = make(chan struct{})

func main() {
    fmt.Println("Hello, World!")
    server()    
    <-quit
}

func server() {
    // Listen for incoming connections.
    listener, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    // Close the listener when the application closes.
    defer listener.Close()
    fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

    // Accept a TCP connection
    conn, err := listener.Accept()
    if err != nil {
        panic(err)
    }

    // Setup server side of yamux
    session, err := yamux.Server(conn, nil)
    if err != nil {
        panic(err)
    }

    
        // Accept a stream
        stream, err := session.AcceptStream()
        if err != nil {
            panic(err)
        }

        go handleStream1(stream)
        // Accept a stream
        stream2, err := session.AcceptStream()
        if err != nil {
            panic(err)
        }

        go handleStream2(stream2)
    
}

func handleStream1(stream *yamux.Stream) {
    fmt.Println(stream.StreamID())    
    const size = 10000 * 1024
    for {
        // Listen for a message
        buf := make([]byte, size)
        n, err := stream.Read(buf)
        if n == 0 {
            panic(err)
        }        
        if err != nil {
            panic(err)
        }        
        fmt.Println(stream.StreamID(),n)
    }    
}
func handleStream2(stream *yamux.Stream) {
    fmt.Println(stream.StreamID())    
    const size = 4
    for {
        // Listen for a message
        buf := make([]byte, size)
        n, err := stream.Read(buf)
        if n != size {
            panic(err)
        }        
        if err != nil {
            panic(err)
        }        
        fmt.Println(stream.StreamID(),n)
    }    
}