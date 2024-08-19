package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// add as many files as you want and stores all the parsed files within the pointer to a template.
	tpl, err := template.ParseGlob("templates/*.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	// err = tpl.Execute(os.Stdout, nil)
	err = tpl.ExecuteTemplate(os.Stdout, "home.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
