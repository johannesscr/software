package database

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

// init will run at the beginning of start-up.
func init() {
	fmt.Println("init database begin")
	DB = ConnectDB()
	fmt.Println("init database end")
}

func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		log.Fatalln("open", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln("ping", err)
	}
	return db
}
