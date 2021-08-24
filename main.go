package main

import (
	"go-echo/db"
	"go-echo/routes"
)

func main() {
	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
