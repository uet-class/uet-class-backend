package main

import (
	"github.com/uet-class/uet-class-backend/config"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/server"
)

func main() {
	config.Init("develop")
	database.InitPostgres()
	database.InitRedis()
	server.Init()
}
