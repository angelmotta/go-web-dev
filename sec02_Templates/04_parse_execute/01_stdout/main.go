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

	if err = tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}
}
