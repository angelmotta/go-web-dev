package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	myServer, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer myServer.Close()

	for {
		conn, err := myServer.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	// Get scanner from Connection (reader)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text() // Read incoming msg
		fmt.Println("Received: ", line)
		if line == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break // End of HTTP Request Headers
		}
	}
	fmt.Println("Code got here")
	io.WriteString(conn, "I see you connected, Bye!\n")
}
