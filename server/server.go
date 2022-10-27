package server

import (
	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/config"
	"github.com/uet-class/uet-class-backend/controllers"
)

func getHome(c *gin.Context) {
	c.String(200, "This is our homepage")
}

func Init() {
	config := config.GetConfig()

	router := gin.Default()

	router.GET("/", getHome)

	authRouter := router.Group("auth")
	{
		auth := new(controllers.AuthController)
		authRouter.POST("/signup", auth.SignUp)
	}

	userRouter := router.Group("user")
	{
		user := new(controllers.UserController)
		userRouter.GET("/:id", user.GetUser)
	}

	router.Run(config.GetString("SERVER_PORT"))
}
