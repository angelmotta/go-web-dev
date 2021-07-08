package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	sages := map[string]string {
		"India":	"Gandhi",
		"America":	"MLK",
		"Meditate":	"Buddha",
		"Love":		"Jesus",
		"Prophet":	"Muhammad",
	}

	if err := tpl.Execute(os.Stdout, sages); err != nil {
		log.Fatalln()
	}
}
