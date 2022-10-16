package server

import (
	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/config"
)

func Init() {
	config := config.GetConfig()

	router := gin.Default()

	// router.GET("/users", controllers.)

	router.Run(config["port"])
}
