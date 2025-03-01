package main

import (
	"fmt"
	"net"
)

func main() {
	// listen for incoming connections on port 800
	port := 8080

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		fmt.Println(err)
		return
	}

	// ensure listener is closed when program terminates
	defer ln.Close()

	// accept and handle incoming conn
	for {
		conn, err := ln.Accept()

		if err != nil {
			fmt.Println(err)
			continue
		}

		// NOTE: "go" executes as goroutine; lightweight thread managed by go runtime
		// makes sense here, connections should be done async!
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// close conn when done
	// NOTE: defer schedules call to be ran right before enclosing function returns
	defer conn.Close()

	// read incoming data

	// alloc 1024 charlength for msg
	buf := make([]byte, 1024)

	// read buffer from conn
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	// print incoming data!
	fmt.Printf("Received: %s", buf)
}
