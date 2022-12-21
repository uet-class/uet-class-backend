package middlewares

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/controllers"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
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
		controllers.AbortWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	_, err = rdb.Get(ctx, reqSessionId).Result()
	if err != nil {
		if err == redis.Nil {
			controllers.AbortWithError(c, http.StatusUnauthorized, "Session not found")
			return
		}
		controllers.AbortWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.Next()
}

func IsAdmin(c *gin.Context) {
	sessionId, err := c.Cookie("sessionId")
	if err != nil {
		controllers.AbortWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var reqUser *models.User
	if reqUser, err = controllers.GetUserBySessionId(sessionId); err != nil {
		controllers.AbortWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !reqUser.IsAdmin {
		controllers.AbortWithError(c, http.StatusUnauthorized, "You are not administrator")
		return
	}

	c.Next()
}
