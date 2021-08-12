package main

import (
	"bufio"
	"fmt"
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
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// Read request
	method, url := request(conn)

	// Write response
	respond(conn, method, url)
}

func request(conn net.Conn) (string, string) {
	i := 0
	var methodName, url string
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			methodName = strings.Fields(ln)[0]
			fmt.Printf("*** METHOD: %v ***\n", methodName)
			url = strings.Fields(ln)[1]
			fmt.Printf("*** URL Request: %v ***\n", url)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return methodName, url
}

func respond(conn net.Conn, method string, url string) {
	if method == "GET" {
		var body string

		if url == "/" {
			body = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>Hello World nene</body></html>`
		} else {
			body = `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body>Other URL nene</body></html>`
		}

		// Write in the connection
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n") // headers done
		fmt.Fprint(conn, body)
	}

}