package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
	"gorm.io/gorm"
)

type ClassController struct{}

func (class ClassController) CreateClass(c *gin.Context) {
	db := database.GetDatabase()

	sessionId, err := c.Cookie("sessionId")
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}

	reqUser, err := getUserBySessionId(sessionId)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}

	var reqClass models.Class
	if err := c.BindJSON(&reqClass); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}

	reqClass.Teachers = append(reqClass.Teachers, *reqUser)
	if err := db.Create(&reqClass).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	reqUser.ClassTeacher = append(reqUser.ClassTeacher, reqClass)
	if err := db.Save(&reqUser).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (class ClassController) GetUserClasses(c *gin.Context) {
	db := database.GetDatabase()

	sessionId, err := c.Cookie("sessionId")
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}

	reqUser, err := getUserBySessionId(sessionId)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err = db.Model(&reqUser).Preload("ClassTeacher").Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, reqUser)
}

func (class ClassController) GetClass(c *gin.Context) {
	db := database.GetDatabase()

	var matchedClass models.Class
	if err := db.Preload("Teachers").First(&matchedClass, c.Param("id")).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseHandler(c, http.StatusNotFound, err)
			return
		}
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, matchedClass)
}

func (class ClassController) DeleteClass(c *gin.Context) {
	db := database.GetDatabase()

	var matchedClass models.Class

	if err := db.Delete(&matchedClass, c.Param("id")).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ResponseHandler(c, http.StatusNotFound, err)
			return
		}
		ResponseHandler(c, http.StatusInternalServerError, err)
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}
