package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		fmt.Println("Before Accept")
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		handle(conn)
	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
	}
	defer conn.Close()
	// We never reach this line until we finish the connection
	fmt.Println("Code got here. Bye!")
}