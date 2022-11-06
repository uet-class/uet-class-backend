package controllers

import (
	"errors"
	"time"

	"github.com/google/uuid"

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
		InternalServerErrorHandler(err, c)
		return
	}

	if err := db.Where(&models.User{Email: reqUser.Email}).First(&matchedUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if reqUser.Password, err = hashPassword(reqUser.Password); err != nil {
			InternalServerErrorHandler(err, c)
			return
		}
		if err := db.Create(&reqUser).Error; err != nil {
			InternalServerErrorHandler(err, c)
			return
		}
	} else {
		MessageHandler("User already exists", c)
		return
	}

	MessageHandler("Success", c)
}

func (auth AuthController) SignIn(c *gin.Context) {
	db := database.GetDatabase()
	rdb := database.GetRedis()

	var reqUser models.User
	var matchedUser models.User

	if err := c.BindJSON(&reqUser); err != nil {
		InternalServerErrorHandler(err, c)
		return
	}

	if err := db.Where(&models.User{Email: reqUser.Email}).First(&matchedUser).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if checkPasswordHash(reqUser.Password, matchedUser.Password) {
			sessionId := uuid.NewString()

			sessionDuration, err := time.ParseDuration("3h")
			if err != nil {
				InternalServerErrorHandler(err, c)
			}

			c.SetCookie("sessionId", sessionId, int(sessionDuration), "/", "uetclass-dev.duckdns.org", false, true)

			err = rdb.Set(database.GetRedisContext(), sessionId, "authorized", sessionDuration).Err()
			if err != nil {
				InternalServerErrorHandler(err, c)
				return
			}

			MessageHandler("Sign in successfully", c)
			return
		}
	}
	MessageHandler("Username or password is incorrect", c)
}

func (auth AuthController) SignOut(c *gin.Context) {
	var reqUser models.User

	if err := c.BindJSON(&reqUser); err != nil {
		InternalServerErrorHandler(err, c)
		return
	}

	rdb := database.GetRedis()

	result, err := rdb.Get(c, reqUser.Email).Result()
	if err != nil {
		InternalServerErrorHandler(err, c)
	}

	print(result)
	MessageHandler("Succeed", c)
}
