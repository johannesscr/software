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
	err := tpl.ExecuteTemplate(os.Stdout, "home.gohtml", 12)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "variable.gohtml", "my variable")
	if err != nil {
		log.Fatalln(err)
	}

	d := struct {
		Slice []string
		Map   map[string]string
	}{
		Slice: []string{"shimano", "sram"},
		Map: map[string]string{
			"scott":   "addict",
			"colnago": "v4rs",
		},
	}

	err = tpl.ExecuteTemplate(os.Stdout, "composite-data.gohtml", d)
	if err != nil {
		log.Fatalln(err)
	}
}
