package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	// Client dial to the TCP Server
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln("Dial error", err)
	}
	defer conn.Close()

	// I read all bytes[] received from the TCP Server
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(bs))
}
