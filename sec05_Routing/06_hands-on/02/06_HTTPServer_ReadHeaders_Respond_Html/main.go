package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

	// Read request from the client
	method, uri := readRequest(conn)
	// Write back response to the client
	writeRespond(conn, method, uri)
}

func readRequest(conn net.Conn) (string, string) {
	var requestMethod, uri string
	myScanner := bufio.NewScanner(conn)

	i := 0
	for myScanner.Scan() {
		line := myScanner.Text()
		if i == 0 {
			requestMethod = strings.Fields(line)[0]
			fmt.Printf("** Method: %v **\n", requestMethod)
			uri = strings.Fields(line)[1]
			fmt.Printf("** URI requested: %v **\n", uri)
		}
		if line == "" {
			break // Blank line ~ Headers done
		}
		i++
	}
	return requestMethod, uri
}

func writeRespond(conn net.Conn, reqMethod string, uri string) {
	msg := "HOLY COW THIS IS LOW LEVEL"
	body := fmt.Sprintf("<!DOCTYPE html><html lang=\"en\"><head><meta charset=\"UTF-8\"><title></title></head><body><h1>%v</h1></body></html>",msg)

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n") // Blank line ~ Headers done
	io.WriteString(conn, body)
}