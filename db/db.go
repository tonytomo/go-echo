package db

import (
	"database/sql"
	"go-echo/config"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var err error
var db *sql.DB

func Init() {
	conf := config.GetConfig()

	connectionString := conf.DB_USERNAME + ":" + conf.DB_PASSWORD + "@tcp(" + conf.DB_HOST + ":" + conf.DB_PORT + ")/" + conf.DB_NAME

	db, err = sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateCon() *sql.DB {
	return db
}
