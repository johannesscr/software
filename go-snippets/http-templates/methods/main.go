package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	fm := template.FuncMap{
		"double": func(x int) int {
			return 2 * x
		},
		"uc": strings.ToUpper,
	}
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("index.gohtml"))
}

type Person struct {
	Name string
}

func (p Person) Gamble() int {
	return rand.Int()
}

func (p Person) Square(x int) int {
	return x * x
}

func main() {
	p := Person{"Frikkie"}
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", p)
	if err != nil {
		log.Fatalln(err)
	}
}
