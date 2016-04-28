package main

import (
    "fmt"
    "net"
    "time"
)

// NewClient is called when new client connected.
func NewClient(connection net.Conn) {
	buf := make([]byte, 512)
    for {
        n, err := connection.Read(buf)
        if nerr, ok := err.(net.Error); ok && !nerr.Timeout() {
            fmt.Printf("Read Error: %q\n", err)
            return
        }
        if n > 0 {
            n, err = connection.Write(buf[:n])
            if nerr, ok := err.(net.Error); ok && !nerr.Timeout() {
                fmt.Printf("Write Error: %q\n", err)
                return
            }
        }
        time.Sleep(50 * time.Millisecond)
    }
}

func main() {	

	listener, _ := net.Listen("tcp", ":11000")   

	for {
		conn, _ := listener.Accept()
		go NewClient(conn)
	}
}