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

func (u UserController) GetUser(c *gin.Context) {
	db := database.GetDatabase()

	var matchedUser models.User
	if err := db.First(&matchedUser, c.Param("id")).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseHandler(c, http.StatusNotFound, err)
			return
		}
		ResponseHandler(c, http.StatusInternalServerError, err)
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
