package main

import (
	"github.com/uet-class/uet-class-backend/config"
	// "github.com/uet-class/uet-class-backend/db"
	"github.com/uet-class/uet-class-backend/server"
)

func main() {
	config.Init()
	// db.Init()
	server.Init()
}
