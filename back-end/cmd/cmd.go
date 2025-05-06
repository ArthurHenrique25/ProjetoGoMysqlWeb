package main

import (
	"projeto/back-end/server"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	server.StartServer()

}
