package middlewares

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/controllers"
	"github.com/uet-class/uet-class-backend/database"
)

var ctx context.Context = context.Background()

func AuthRequired(c *gin.Context) {
	rdb := database.GetRedis()

	reqSessionId, err := c.Cookie("sessionId")
	if err != nil {
		if err == http.ErrNoCookie {
			controllers.AbortWithError(c, http.StatusUnauthorized, "Cookie not found")
			return
		}
		controllers.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}

	_, err = rdb.Get(ctx, reqSessionId).Result()
	if err != nil {
		if err == redis.Nil {
			controllers.AbortWithError(c, http.StatusUnauthorized, "Session not found")
			return
		}
		controllers.AbortWithError(c, http.StatusInternalServerError, err)
		return
	}
	c.Next()
}
