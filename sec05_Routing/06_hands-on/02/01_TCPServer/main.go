package main

import (
	"io"
	"log"
	"net"
)

func main() {
	myServer, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panicln(err)
	}
	defer myServer.Close()

	for {
		conn, err := myServer.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
		// be careful close the connection after the task is done
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "I see you connected bro! Bye\n")
}
