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

func getUserByUserId(userId string) (*models.User, error) {
	db := database.GetDatabase()

	matchedUser := new(models.User)
	if err := db.First(&matchedUser, userId).Error; err != nil {
		return nil, err
	}
	return matchedUser, nil
}

func getUserByUserEmail(userEmail string) (*models.User, error) {
	db := database.GetDatabase()

	matchedUser := new(models.User)
	if err := db.Where(&models.User{Email: userEmail}).First(&matchedUser).Error; err != nil {
		return nil, err
	}
	return matchedUser, nil
}

func GetUserBySessionId(sessionId string) (*models.User, error) {
	userId, err := getUserIdBySessionId(sessionId)
	if err != nil {
		return nil, err
	}

	matchedUser, err := getUserByUserId(userId)
	if err != nil {
		return nil, err
	}
	return matchedUser, nil
}

func (u UserController) GetUser(c *gin.Context) {
	matchUser, err := getUserByUserId(c.Param("id"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, matchUser)
}

func (u UserController) DeleteUser(c *gin.Context) {
	db := database.GetDatabase()

	var matchedUser models.User
	if err := db.Where(&models.User{Email: c.Param("email")}).First(&matchedUser).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Delete(&matchedUser).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (u UserController) UpdateUser(c *gin.Context) {
	db := database.GetDatabase()

	var updatedUser models.User
	if err := c.BindJSON(&updatedUser); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	matchedUser, err := getUserByUserId(c.Param("id"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	matchedUser.AvatarUrl = updatedUser.AvatarUrl
	matchedUser.FullName = updatedUser.FullName
	matchedUser.DateOfBirth = updatedUser.DateOfBirth
	matchedUser.PhoneNumber = updatedUser.PhoneNumber
	matchedUser.Password, err = hashPassword(updatedUser.Password)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Save(&matchedUser).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, matchedUser)
}
