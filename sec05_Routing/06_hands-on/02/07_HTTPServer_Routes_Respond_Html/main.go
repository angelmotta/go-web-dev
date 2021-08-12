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
	switch {
	case reqMethod == "GET" && uri == "/":
		handleIndex(conn)
	case reqMethod == "GET" && uri == "/apply":
		handleApply(conn)
	case reqMethod == "POST" && uri == "/apply":
		handleApplyPost(conn)
	default:
		handleDefault(conn)
	}
}

func handleIndex(conn net.Conn) {
	body :=
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET INDEX</title>
		</head>
		<body>
			<h1>"GET INDEX"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>
		`
	sendResponse(conn, body)
}

func handleApply(conn net.Conn) {
	body :=
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>GET DOG</title>
		</head>
		<body>
			<h1>"GET APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
			<form action="/apply" method="POST">
			<input type="hidden" value="In my good death">
			<input type="submit" value="submit">
			</form>
		</body>
		</html>
		`
	sendResponse(conn, body)
}

func handleApplyPost(conn net.Conn) {
	body :=
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>POST APPLY</title>
		</head>
		<body>
			<h1>"POST APPLY"</h1>
			<a href="/">index</a><br>
			<a href="/apply">apply</a><br>
		</body>
		</html>
		`

	sendResponse(conn, body)
}

func handleDefault(conn net.Conn) {
	body :=
		`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>default</title>
		</head>
		<body>
			<h1>"default"</h1>
		</body>
		</html>
		`

	sendResponse(conn, body)
}

func sendResponse(c net.Conn, body string) {
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n") // Blank line ~ End Header
	io.WriteString(c, body)  // Body of response
}