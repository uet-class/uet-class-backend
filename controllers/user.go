package controllers

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/config"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
)

const GoogleCloudStoragePrefix = "https://storage.googleapis.com"

type UserController struct{}

func infoIsChanged(newInfo string, oldInfo string) bool {
	if newInfo != oldInfo && newInfo != "" {
		return true
	}
	return false
}

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

func getUserBySessionId(sessionId string) (*models.User, error) {
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

func uploadFile(bucketName string, fileName string, file multipart.FileHeader) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	uploadFile, err := file.Open()
	if err != nil {
		return err
	}
	defer uploadFile.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	objectHandle := client.Bucket(bucketName).Object(fileName)
	objectHandle = objectHandle.If(storage.Conditions{DoesNotExist: true})

	objectWriter := objectHandle.NewWriter(ctx)
	if _, err = io.Copy(objectWriter, uploadFile); err != nil {
		return err
	}
	if err := objectWriter.Close(); err != nil {
		return err
	}
	return nil
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

	if updatedUser.Password != "" && !checkPasswordHash(updatedUser.Password, matchedUser.Password) {
		if matchedUser.Password, err = hashPassword(updatedUser.Password); err != nil {
			ResponseHandler(c, http.StatusInternalServerError, err.Error())
			return
		}
	}
	
	if infoIsChanged(updatedUser.FullName, matchedUser.FullName) {
		matchedUser.FullName = updatedUser.FullName
	}

	if infoIsChanged(updatedUser.DateOfBirth, matchedUser.DateOfBirth) {
		matchedUser.DateOfBirth = updatedUser.DateOfBirth
	}

	if infoIsChanged(updatedUser.PhoneNumber, matchedUser.PhoneNumber) {
		matchedUser.PhoneNumber = updatedUser.PhoneNumber
	}
	// matchedUser.AvatarUrl = updatedUser.AvatarUrl

	if infoIsChanged(updatedUser.FullName, matchedUser.FullName) {
		matchedUser.FullName = updatedUser.FullName
	}

	if infoIsChanged(updatedUser.DateOfBirth, matchedUser.DateOfBirth) {
		matchedUser.DateOfBirth = updatedUser.DateOfBirth
	}

	if infoIsChanged(updatedUser.PhoneNumber, matchedUser.PhoneNumber) {
		matchedUser.PhoneNumber = updatedUser.PhoneNumber
	}
	// matchedUser.AvatarUrl = updatedUser.AvatarUrl

	if err := db.Save(&matchedUser).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, matchedUser)
}

func (u UserController) UploadUserAvatar(c *gin.Context) {
	db := database.GetDatabase()
	conf := config.GetConfig()
	bucketName := conf.GetString("GCS_BUCKET_AVATARS")

	avatarImage, err := c.FormFile("avatar")
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := uploadFile(bucketName, avatarImage.Filename, *avatarImage); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}

	avatarUrl := fmt.Sprintf("%s/%s/%s", GoogleCloudStoragePrefix, bucketName, avatarImage.Filename)

	matchedUser, err := getUserByUserId(c.Param("id"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	matchedUser.AvatarUrl = avatarUrl
	if err := db.Save(&matchedUser).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	response := map[string]interface{}{"avatarUrl": avatarUrl}
	ResponseHandler(c, http.StatusOK, response)
}
