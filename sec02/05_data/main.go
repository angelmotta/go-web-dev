package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	// Using Must() helper function to check for errors
	tpl = template.Must(template.ParseGlob("tpl.gohtml"))
}

func main() {
	// Passing data to execute the template
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
