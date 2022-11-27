package controllers

import (
	"errors"
	"net/http"
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
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}

	if err := db.Where(&models.User{Email: reqUser.Email}).First(&matchedUser).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		if reqUser.Password, err = hashPassword(reqUser.Password); err != nil {
			ResponseHandler(c, http.StatusUnauthorized, err)
			return
		}
		if err := db.Create(&reqUser).Error; err != nil {
			ResponseHandler(c, http.StatusInternalServerError, err)
			return
		}
	} else {
		ResponseHandler(c, http.StatusConflict, "User already exists")
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (auth AuthController) SignIn(c *gin.Context) {
	db := database.GetDatabase()
	rdb := database.GetRedis()

	var reqUser models.User
	var matchedUser models.User

	if err := c.BindJSON(&reqUser); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}

	// Check if user exists
	if err := db.Where(&models.User{Email: reqUser.Email}).First(&matchedUser).Error; !errors.Is(err, gorm.ErrRecordNotFound) {
		if checkPasswordHash(reqUser.Password, matchedUser.Password) {
			// Generate new session
			sessionId := uuid.NewString()
			sessionDuration, err := time.ParseDuration("3h")
			if err != nil {
				ResponseHandler(c, http.StatusInternalServerError, err)
				return
			}
			// Store and send the session cookie back to the client
			c.SetCookie("sessionId", sessionId, int(sessionDuration), "/", "uetclass-dev.duckdns.org", false, true)
			err = rdb.Set(database.GetRedisContext(), sessionId, matchedUser.ID, sessionDuration).Err()
			if err != nil {
				ResponseHandler(c, http.StatusInternalServerError, err)
				return
			}
			response := map[string]interface{}{"sessionId": sessionId, "userId": matchedUser.ID}
			ResponseHandler(c, http.StatusOK, response)
			return
		}
	}
	ResponseHandler(c, http.StatusForbidden, "Username or password is incorrect")
}

func (auth AuthController) SignOut(c *gin.Context) {
	rdb := database.GetRedis()

	reqSessionId, err := c.Cookie("sessionId")
	if err != nil {
		if err == http.ErrNoCookie {
			ResponseHandler(c, http.StatusUnauthorized, err)
			return
		}
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}

	_, err = rdb.Del(database.GetRedisContext(), reqSessionId).Result()
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}
