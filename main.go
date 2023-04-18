package main

import (
	"github.com/kimMichel/webecom-go-api/database"
	"github.com/kimMichel/webecom-go-api/server"
)

func main() {
	database.ConnectDb()

	server := server.NewServer()
	server.Run()
}
