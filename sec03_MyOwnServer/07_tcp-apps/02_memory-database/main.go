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
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// Show instructions (write these instructions in connection)
	io.WriteString(conn, "\r\n*** In-Memory Database ***\r\n\r\n" +
		"USE:\r\n" +
		"\tSET key value \r\n" +
		"\tGET key \r\n" +
		"\tDEL key \r\n\r\n" +
		"EXAMPLE:\r\n" +
		"\tSET fav chocolate \r\n" +
		"\tGET fav \r\n\r\n\r\n")

	// Receive read & write
	data := make(map[string]string)

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)

		if len(fs) < 1 {
			continue
		}

		switch fs[0] {
		case "GET":
			key := fs[1]
			val := data[key]
			fmt.Fprintf(conn, "%s\r\n", val)
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "Expected 3 fields for SET instruction")
				continue
			}
			key := fs[1]
			val := fs[2]
			data[key] = val
		case "DEL":
			key := fs[1]
			delete(data, key)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND " + fs[0] + "\r\n")
			continue
		}
	}
}