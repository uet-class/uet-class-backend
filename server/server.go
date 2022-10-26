package server

import (
	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/config"
)

func getUsers(c *gin.Context) {
	c.String(200, "This is user collection")
}

func getHome(c *gin.Context) {
	c.String(200, "This is our homepage")
}

func Init() {
	config := config.GetConfig()

	router := gin.Default()

	router.GET("/users", getUsers)
	router.GET("/", getHome)

	router.Run(config.GetString("SERVER_PORT"))
}
