package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ybuilds/slash/utils"
)

var DB *sql.DB

func init() {
	var err error

	var (
		DB_USER = utils.GetValue("DB_USER")
		DB_PASS = utils.GetValue("DB_PASS")
		DB_HOST = utils.GetValue("DB_HOST")
		DB_NAME = utils.GetValue("DB_NAME")
	)

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", DB_USER, DB_PASS, DB_HOST, DB_NAME)

	DB, err = sql.Open("mysql", connectionString)
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
