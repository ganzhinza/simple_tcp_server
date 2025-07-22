package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatalf("failed to connect server:", err)
	}

	buff := make([]byte, 10)

	bytesRead, err := conn.Read(buff)
	if err != nil {
		log.Fatalf("failed to read from connection:", err)
	}

	if bytesRead != 3 && string(buff[:4]) != "OK\n" {
		log.Fatalf("wrong responce")
	}

	fmt.Println(string(buff))
}
