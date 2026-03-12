package database

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func init() {
	var err error

	DB, err = sql.Open("mysql", "root:kjm40438@tcp(localhost:3306)/slash")
	if err != nil {
		log.Fatalln("error creating database connection: ", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Println("error pinging database: ", err)
		return
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
