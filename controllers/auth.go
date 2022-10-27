package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
	"gorm.io/gorm"
)

type AuthController struct{}

func (auth AuthController) SignUp(c *gin.Context) {
	db := database.GetDatabase()

	var reqUser models.User
	var matchedUser models.User

	if err := c.BindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Where(&models.User{Email: reqUser.Email}).First(&matchedUser).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		if err := db.Create(&reqUser).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "User already exists!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": reqUser})
}

func (auth AuthController) SignIn(c *gin.Context) {
	session := sessions.Default(c)
	db := database.GetDatabase()

	var reqUser models.User
	var matchedUser models.User

	if err := c.BindJSON(&reqUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := db.Where(&models.User{Email: reqUser.Email}).First(&matchedUser).Error

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if matchedUser.Password == reqUser.Password {
			session.Set("authorized", true)
			if err := session.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			c.JSON(http.StatusOK, gin.H{"Is authorized": true})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"error": "Username or password is incorrect"})
}

func (auth AuthController) SignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()

	c.JSON(http.StatusOK, session.Get("authorized"))
}
