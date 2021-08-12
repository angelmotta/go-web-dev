package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	myName := "Angel Motta"
	strPage := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Templates</title>
	<body>
	<h1>` + myName + `</h1>
	</body>
	</html>`)

	myFile, err := os.Create("index.html")
	if err != nil {
		log.Fatalln("Error creating html file", err)
	}
	defer myFile.Close()

	if _, err := io.Copy(myFile, strings.NewReader(strPage)); err != nil {
		log.Fatalln("Error writing html file", err)
	}
}
