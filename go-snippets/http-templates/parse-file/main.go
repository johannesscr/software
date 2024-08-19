package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// add as many files as you want and stores all the parsed files within the
	// pointer to a template.
	tpl, err := template.ParseFiles("templates/index.gohtml", "templates/home.gohtml")
	// after the files are parsed, they can all be executed using the
	// ExecuteTemplate method using only the filename.
	if err != nil {
		log.Fatalln(err)
	}
	err = tpl.ExecuteTemplate(os.Stdout, "home.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
