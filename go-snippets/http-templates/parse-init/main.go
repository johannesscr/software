package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

// Set up the parsing of template before the program runs
func init() {
	// The `Must` function helps and does the error handling for use, it just
	// panics if the error is non-nil.
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "home.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
