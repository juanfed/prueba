package main

import (
	_ "github.com/go-sql-driver/mysql"
	"prueba/server"
)

func main() {
	server.Start()
}
