package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

// Set up the parsing of template before the program runs
func init() {
	fm := template.FuncMap{
		"uc": strings.ToUpper,
	}
	// The `Must` function helps and does the error handling for use, it just
	// panics if the error is non-nil.
	//
	// Create a new template, Then add all the functions in the func map.
	// Only now can the files be parsed, else it will not be able to parse all
	// functions used.
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "home.gohtml", "enter as lower case")
	if err != nil {
		log.Fatalln(err)
	}
}
