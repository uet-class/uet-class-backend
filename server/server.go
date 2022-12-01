package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/config"
	"github.com/uet-class/uet-class-backend/controllers"
	"github.com/uet-class/uet-class-backend/middlewares"
)

func getHome(c *gin.Context) {
	c.String(200, "These are our APIs.")
}

func Init() {
	router := gin.Default()
	config := config.GetConfig()

	corsPolicy := cors.DefaultConfig()
	corsPolicy.AllowOrigins = []string{
		"http://localhost:3000",
		"http://uc-frontend",
		"https://uetclass.duckdns.org",
	}
	corsPolicy.AllowCredentials = true
	corsPolicy.ExposeHeaders = []string{"Content-Length"}
	corsPolicy.AllowMethods = []string{"GET", "POST", "DELETE"}

	router.Use(cors.New(corsPolicy))
	router.Use(sessions.Sessions("uc-session", cookie.NewStore([]byte("SessionSecret"))))

	router.GET("/api", getHome)

	authRouter := router.Group("api/auth")
	{
		auth := new(controllers.AuthController)
		authRouter.POST("/signup", auth.SignUp)
		authRouter.POST("/signin", auth.SignIn)
		authRouter.POST("/signout", middlewares.AuthRequired, auth.SignOut)
	}

	userRouter := router.Group("api/user").Use(middlewares.AuthRequired)
	{
		user := new(controllers.UserController)
		userRouter.GET("/:id", user.GetUser)
		userRouter.POST("/:id/", user.UpdateUser)
		userRouter.POST("/:id/upload-avatar", user.UploadUserAvatar)
		userRouter.DELETE("/:email", user.DeleteUser)
	}

	classRouter := router.Group("api/class").Use(middlewares.AuthRequired)
	{
		class := new(controllers.ClassController)
		classRouter.POST("", class.CreateClass)
		classRouter.POST("/:id/add-student", class.AddStudent)
		classRouter.POST("/:id/send-invitation", class.SendInvitation)
		classRouter.GET("/accept-invitation", class.AcceptInvitation)
		classRouter.GET("/all", class.GetUserClasses)
		classRouter.GET("/:id", class.GetClass)
		classRouter.DELETE("/:id", class.DeleteClass)
	}

	reportRouter := router.Group("api/report").Use(middlewares.AuthRequired)
	{
		report := new(controllers.ReportController)
		reportRouter.POST("", report.CreateReport)
		reportRouter.GET("", report.GetUserReports)
	}

	router.Run(config.GetString("SERVER_PORT"))
}
