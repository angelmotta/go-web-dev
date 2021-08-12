package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	myNewHtmlFile, err := os.Create("index.html")
	if err != nil {
		log.Println("Error creating new html file", err)
	}
	defer myNewHtmlFile.Close()

	if err = tpl.Execute(myNewHtmlFile, nil); err != nil {
		log.Fatalln("Error in Execute() template", err)
	}
}
