package main

import "fmt"
import "sample/yamux"
import "net"
import "time"
//import "os"

const (
    CONN_HOST = "localhost"
    CONN_PORT = "3333"
    CONN_TYPE = "tcp"
)

var quit = make(chan struct{})

func main() {
    fmt.Println("Hello, World!")    
    client()
    <-quit
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
        stream1, err := session.OpenStream()
        if err != nil {
            panic(err)
        }
        go handleStream(stream1)
    
                // Open a new stream
                stream2, err := session.OpenStream()
                if err != nil {
                    panic(err)
                }
                go handleStream2(stream2)

    

}

func handleStream(stream *yamux.Stream) {
    fmt.Println(stream.StreamID())    
    data := make([]byte, 100000*1024)
    for idx := range data {
        data[idx] = byte(idx % 256)
    }        

    for {
        // Stream implements net.Conn
        stream.Write(data)
        time.Sleep(1 * time.Second)    
        }    

}
func handleStream2(stream *yamux.Stream) {
    fmt.Println(stream.StreamID())    
    for {
        // Stream implements net.Conn
        stream.Write([]byte("ping"))
        time.Sleep(100 * time.Millisecond)    
        }    
}