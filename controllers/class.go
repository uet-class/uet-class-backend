package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uet-class/uet-class-backend/config"
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

func getRecipientsWithInvitationIndex(c *gin.Context) (map[string]string, error) {
	rdb := database.GetRedis()

	recipients := make(map[string]string)
	var recipientEmails []string
	if err := c.BindJSON(&recipientEmails); err != nil {
		return nil, err
	}

	for _, recipientEmail := range recipientEmails {
		invitationIndex := uuid.NewString()
		invitationDuration, err := time.ParseDuration("24h")
		if err != nil {
			return nil, err
		}
		err = rdb.Set(database.GetRedisContext(), invitationIndex, recipientEmail, invitationDuration).Err()
		if err != nil {
			return nil, err
		}
		recipients[invitationIndex] = recipientEmail
	}
	return recipients, nil
}

func (class ClassController) SendInvitationEmail(c *gin.Context) {
	envConfig := config.GetConfig()

	recipients, err := getRecipientsWithInvitationIndex(c)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	sender := envConfig.GetString("SMTP_EMAIL_USERNAME")
	hostname := "smtp.gmail.com"
	smtpAddress := fmt.Sprintf("%s:%s", hostname, envConfig.GetString("SMTP_PORT"))
	auth := smtp.PlainAuth("", sender, envConfig.GetString("SMTP_EMAIL_PASSWORD"), hostname)

	for invitationIndex, recipientEmail := range recipients {
		fmt.Println(invitationIndex)

		recipientHeader := fmt.Sprintf("To: %s\r\n", recipientEmail)
		subjectHeader := "Subject: Invitation to a new class!\r\n"
		body := fmt.Sprintf("Confirmation link: %s\r\n", invitationIndex)

		message := []byte(recipientHeader + subjectHeader + "\r\n" + body)
		if err := smtp.SendMail(smtpAddress, auth, sender, []string{recipientEmail}, message); err != nil {
			ResponseHandler(c, http.StatusInternalServerError, err)
			return
		}
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}
