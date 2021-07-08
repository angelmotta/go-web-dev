package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	tpl, err := template.ParseFiles("one.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// using the pointer to template 'tpl'
	if err = tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}

	// Load another file into template
	tpl, err = tpl.ParseFiles("two.txt", "vespa.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// Execute specific template
	if err = tpl.ExecuteTemplate(os.Stdout, "vespa.txt", nil); err != nil {
		log.Fatalln(err)
	}

	if err = tpl.ExecuteTemplate(os.Stdout, "two.txt", nil); err != nil {
		log.Fatalln(err)
	}

	if err = tpl.ExecuteTemplate(os.Stdout, "one.txt", nil); err != nil {
		log.Fatalln("Error here:",err)
	}

	// Execute all files inside template
	fmt.Println("**** Execute all templates ****")
	if err = tpl.Execute(os.Stdout, nil); err != nil {
		log.Fatalln(err)
	}
}
