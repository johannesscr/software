package main

import (
	"log"
	"os"
	"text/template"
)

func main() {
	// we provide a template name and then parse the template content
	// "Must" does the error handling.
	tpl := template.Must(template.New("some-template").Parse(`this is the template content with parsed data: {{.}}`))
	err := tpl.ExecuteTemplate(os.Stdout, "some-template", "some data")
	if err != nil {
		log.Fatalln(err)
	}
}
