package main

import "fmt"

func main() {
	myName := "Angel Motta"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Templates</title>
	<body>
	<h1>` + myName + `</h1>
	</body>`

	fmt.Println(tpl)
}
