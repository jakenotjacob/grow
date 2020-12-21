package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) //connection aborted
			continue
		}
		if conn != nil {
			fmt.Println("Client connected!")
			handleConn(conn) // handle one connection at a time
		}
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:05:05\n"))
		if err != nil {
			fmt.Println("Client disconnected!")
			return //client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
