package main

import (
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/server"
	"github.com/uet-class/uet-class-backend/storage"
)

func main() {
	storage.InitStorageClient()
	database.InitPostgres()
	database.InitRedis()
	server.Init()
}
