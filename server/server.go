package server

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/config"
	"github.com/uet-class/uet-class-backend/controllers"
	"github.com/uet-class/uet-class-backend/middlewares"
	"github.com/gin-contrib/cors"
)

func getHome(c *gin.Context) {
	c.String(200, "This is our homepage")
}

func Init() {
	config := config.GetConfig()

	router := gin.Default()
  router.Use(cors.Default())
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

	classRouter := router.Group("class").Use(middlewares.AuthRequired)
	{
		class := new(controllers.ClassController)
		classRouter.POST("/", class.CreateClass)
		classRouter.GET("/:id", class.GetClass)
		classRouter.DELETE("/:id", class.DeleteClass)
	}

	reportRouter := router.Group("report").Use(middlewares.AuthRequired)
	{
		report := new(controllers.ReportController)
		reportRouter.POST("/", report.CreateReport)
		reportRouter.GET("/", report.GetUserReports)
	}

	router.Run(config.GetString("SERVER_PORT"))
}
