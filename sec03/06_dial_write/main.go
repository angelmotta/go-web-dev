package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// Client code
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	// Write to connection
	fmt.Fprintln(conn, "client: I dialed you (write action)")
}
