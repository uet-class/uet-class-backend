package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/controllers"
	"github.com/uet-class/uet-class-backend/database"
)

var ctx context.Context = context.Background()

func AuthRequired(c *gin.Context) {
	fmt.Println("authorizing")
	rdb := database.GetRedis()

	reqSessionId, err := c.Cookie("sessionId")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Println("no cookies")
			controllers.UnauthorizedErrorHandler(err, c)
			return
		}
		fmt.Println("server is dead")
		controllers.InternalServerErrorHandler(err, c)
		return
	}

	fmt.Println(reqSessionId)

	isAuthorized, err := rdb.Get(ctx, reqSessionId).Result()
	if err != nil {
		if err == redis.Nil {
			controllers.UnauthorizedErrorHandler(err, c)
			return
		}
		controllers.InternalServerErrorHandler(err, c)
		return
	}

	fmt.Println(isAuthorized)
	c.Next()
}
