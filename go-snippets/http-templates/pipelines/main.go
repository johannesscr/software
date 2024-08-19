package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var tpl *template.Template

func formatDate(t time.Time) string {
	return t.Format("2006-01-02")
}

var fm = template.FuncMap{
	"fmtDateYMD": formatDate,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("templates/*.gohtml"))
}

func main() {
	data := struct {
		MyDate time.Time
	}{
		time.Now(),
	}
	err := tpl.ExecuteTemplate(os.Stdout, "home.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
