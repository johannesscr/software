package main

import (
	"fmt"
	"github.com/johannesscr/software/database/sqlite/database"
	"log"
)

func init() {
	fmt.Println("begin init main")
}

func main() {
	fmt.Println("hello")
	err := database.DB.Ping()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("end")
}
