package main

import (
	"github.com/kimMichel/webecom-go-api/database"
	"github.com/kimMichel/webecom-go-api/initializers"
	"github.com/kimMichel/webecom-go-api/server"
)

func init() {
	initializers.LoadEnv()
}

func main() {
	database.ConnectDb()

	server := server.NewServer()
	server.Run()
}
