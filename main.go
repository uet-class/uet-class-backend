package main

import (
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/gcs"
	"github.com/uet-class/uet-class-backend/server"
)

func main() {
	gcs.InitStorageClient()
	database.InitPostgres()
	database.InitRedis()
	server.Init()
}
