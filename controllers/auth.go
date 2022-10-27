package controllers

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct{}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (auth AuthController) SignUp(c *gin.Context) {
	db := database.GetDatabase()

	var reqUser models.User
	var matchedUser models.User

	if err := c.BindJSON(&reqUser); err != nil {
		ErrorHandler(err, c)
		return
	}

	if err := db.Where(&models.User{Email: reqUser.Email}).First(&matchedUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if reqUser.Password, err = hashPassword(reqUser.Password); err != nil {
			ErrorHandler(err, c)
			return
		}
		if err := db.Create(&reqUser).Error; err != nil {
			ErrorHandler(err, c)
			return
		}
	} else {
		MessageHandler("User is already exists", c)
		return
	}

	MessageHandler("Success", c)
}

func (auth AuthController) SignIn(c *gin.Context) {
	session := sessions.Default(c)
	db := database.GetDatabase()

	var reqUser models.User
	var matchedUser models.User

	if err := c.BindJSON(&reqUser); err != nil {
		ErrorHandler(err, c)
		return
	}

	if err := db.Where(&models.User{Email: reqUser.Email}).First(&matchedUser).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if checkPasswordHash(reqUser.Password, matchedUser.Password) {
			session.Set("authorized", true)
			if err := session.Save(); err != nil {
				ErrorHandler(err, c)
				return
			}
			MessageHandler("Success", c)
			return
		}
	}

	MessageHandler("Username or password is incorrect", c)
}

func (auth AuthController) SignOut(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	MessageHandler(session.Get("authorized"), c)
}
