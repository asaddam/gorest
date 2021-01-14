package main

import (
	"github.com/gorest1/db"
	"github.com/gorest1/routes"
	
)

func main() {
	db.Init()
	
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1000"))
}