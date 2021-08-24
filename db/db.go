package db

import (
	"database/sql"
	"fmt"
	"go-echo/config"
	"log"

	_ "github.com/lib/pq"
)

var err error
var db *sql.DB

func Init() {
	conf := config.GetConfig()

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", conf.DB_USERNAME, conf.DB_PASSWORD, conf.DB_NAME)

	db, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}
}

func CreateCon() *sql.DB {
	return db
}
