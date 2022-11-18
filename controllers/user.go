package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
	"gorm.io/gorm"
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

func getUserByUserId(userId string) (*models.User, error) {
	db := database.GetDatabase()

	matchedUser := new(models.User)
	if err := db.First(&matchedUser, userId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, err
	}
	return matchedUser, nil
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

func (u UserController) GetUser(c *gin.Context) {
	userId := c.Param("id")

	matchedUser, err := getUserByUserId(userId)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, matchedUser)
}

func (u UserController) DeleteUser(c *gin.Context) {
	db := database.GetDatabase()

	var matchedUser models.User

	if err := db.Delete(&matchedUser, c.Param("id")).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseHandler(c, http.StatusNotFound, err)
			return
		}
		ResponseHandler(c, http.StatusInternalServerError, err)
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}
