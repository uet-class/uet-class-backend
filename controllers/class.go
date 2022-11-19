package controllers

import (
	"errors"
	"net/http"
	"net/smtp"

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
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	reqClass.Teachers = append(reqClass.Teachers, *reqUser)
	if err := db.Create(&reqClass).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (class ClassController) AddStudent(c *gin.Context) {
	db := database.GetDatabase()

	var reqUser models.User
	if err := c.BindJSON(&reqUser); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	var matchedClass models.Class
	if err := db.First(&matchedClass, c.Param("id")).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Model(&matchedClass).Association("Students").Append(&reqUser); err != nil {
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

	reqUserId, err := getUserIdBySessionId(sessionId)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	var matchedTeacherClasses []models.Class
	if err = db.Raw("SELECT * FROM classes WHERE id IN (SELECT class_id FROM teacher_class WHERE user_id=?)", reqUserId).Scan(&matchedTeacherClasses).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	var matchedStudentClasses []models.Class
	if err = db.Raw("SELECT * FROM classes WHERE id IN (SELECT class_id FROM student_class WHERE user_id=?)", reqUserId).Scan(&matchedStudentClasses).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	result := map[string][]models.Class{
		"teacherClasses": matchedTeacherClasses,
		"studentClasses": matchedStudentClasses,
	}
	ResponseHandler(c, http.StatusOK, result)
}

func (class ClassController) GetClass(c *gin.Context) {
	db := database.GetDatabase()

	var matchedClass models.Class
	if err := db.Preload("Teachers").First(&matchedClass, c.Param("id")).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
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

func (class ClassController) SendInvitationEmail(c *gin.Context) {
	var (
		from       = "uetclass.notifier@gmail.com"
		msg        = []byte("dummy message")
		recipients = []string{"thainguyen.uet@gmail.com"}
	)
	hostname := "smtp.gmail.com"
	auth := smtp.PlainAuth("", "uetclass.notifier@gmail.com", "xvvpncjpatchxgjp", hostname)

	err := smtp.SendMail(hostname+":25", auth, from, recipients, msg)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}
