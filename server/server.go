package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/config"
	"github.com/uet-class/uet-class-backend/controllers"
	"github.com/uet-class/uet-class-backend/middlewares"
)

func getHome(c *gin.Context) {
	c.String(200, "This is our homepage")
}

func Init() {
	config := config.GetConfig()

	router := gin.Default()
	router.Use(sessions.Sessions("uc-session", cookie.NewStore([]byte("SessionSecret"))))

	router.GET("/", getHome)

	authRouter := router.Group("auth")
	{
		auth := new(controllers.AuthController)
		authRouter.POST("/signup", auth.SignUp)
		authRouter.POST("/signin", auth.SignIn)
		authRouter.POST("/signout", middlewares.AuthRequired, auth.SignOut)
	}

	userRouter := router.Group("user").Use(middlewares.AuthRequired)
	{
		user := new(controllers.UserController)
		userRouter.GET("/:id", user.GetUser)
		userRouter.DELETE("/:id", user.DeleteUser)
	}

	router.Run(config.GetString("SERVER_PORT"))
}
