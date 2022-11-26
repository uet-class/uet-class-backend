package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
)

type UserController struct{}

func getUserIdBySessionId(sessionId string) (string, error) {
	rdb := database.GetRedis()

	userId, err := rdb.Get(database.GetRedisContext(), sessionId).Result()
	if err != nil {
		return "", err
	}
	return userId, nil
}

func getUserBySessionId(sessionId string) (*models.User, error) {
	db := database.GetDatabase()
	rdb := database.GetRedis()

	userId, err := rdb.Get(database.GetRedisContext(), sessionId).Result()
	if err != nil {
		return nil, err
	}

	matchedUser := new(models.User)
	if err := db.First(&matchedUser, userId).Error; err != nil {
		return nil, err
	}
	return matchedUser, nil
}

func (u UserController) GetUserInformation(c *gin.Context) {
	sessionId, err := c.Cookie("sessionId")
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	matchedUser, err := getUserBySessionId(sessionId)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, matchedUser)
}

func (u UserController) DeleteUser(c *gin.Context) {
	db := database.GetDatabase()

	if err := db.Where("email = ?", c.Param("email")).Delete(&models.User{}).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (u UserController) UpdateUser(c *gin.Context) {
	
}
