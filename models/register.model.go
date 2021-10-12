package models

import (
	"fmt"

	"go-echo/db"
	"go-echo/helpers"
)

func Register(username, password string) (bool, error) {
	hashedPwd, _ := helpers.HashPassword(password)

	con := db.CreateCon()

	sqlStatement := "INSERT INTO users (username, password) VALUES (?, ?)"

	_, err := con.Query(sqlStatement, username, hashedPwd)

	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
