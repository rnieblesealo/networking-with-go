package main

import (
	"fmt"
	"net"
)

func main() {
	// connect to server via tcp using dial func
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}

	// send data to svr
	_, err = conn.Write([]byte("Hello, server!"))
	if err != nil {
		fmt.Println(err)
		return
	}

	conn.Close()
}
