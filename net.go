package main

import (
    "net"
    "log"
    "fmt"
    "os"
    "io"
)

func main() {
    ln, err := net.Listen("tcp", ":8080")

    if err != nil {
    	log.Fatal(err)
    }

    fmt.Println("Opened on 8080")

    for {
    	conn, err := ln.Accept()
    	if err != nil {
    		log.Fatal(err)
    	}
        fmt.Println("Client connect")
    	go handleConnection(conn)
    }
}

func handleConnection(conn net.Conn) {
    defer func() {
        fmt.Println("Client disconnect")
        conn.Close()
    }()

    fmt.Fprintf(conn, "Hello, world!\n")
    io.Copy(os.Stdout, conn)
}
