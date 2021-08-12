package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	myListener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer myListener.Close()

	for {
		conn, err := myListener.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	// Reader of Connection
	scanner := bufio.NewScanner(conn)
	// Scan Connection
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		// Write to the Connection
		fmt.Fprintf(conn, "I heard you say: %s\n", line)
	}
	defer conn.Close()
	// We never reach this code until connection is closed
	fmt.Println("Good bye nene!")
}
