package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.ListenAndServe(":8080", nil)
}

func dog (w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w,
		`<!-- not serving from our server (no route /toby.jpg) -->
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
			<title>GET INDEX</title>
			</head>
			<body>
			<img src="/toby.jpg">
			</body>
			</html>
			`)
}