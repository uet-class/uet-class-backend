package controllers

import (
	"fmt"
	"net/http"
	"net/smtp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/uet-class/uet-class-backend/config"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
)

type ClassController struct{}

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

func getClassByClassId(id string) (*models.Class, error) {
	db := database.GetDatabase()

	matchedClass := new(models.Class)
	if err := db.First(&matchedClass, id).Error; err != nil {
		return nil, err
	}
	return matchedClass, nil
}

func getUserEmailByInvitationId(invitationId string) (string, error) {
	rdb := database.GetRedis()

	userEmail, err := rdb.Get(database.GetRedisContext(), invitationId).Result()
	if err != nil {
		return "", err
	}
	return userEmail, nil
}

func (class ClassController) CreateClass(c *gin.Context) {
	conf := config.GetConfig()
	db := database.GetDatabase()

	sessionId, err := c.Cookie("sessionId")
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}

	reqUser, err := GetUserBySessionId(sessionId)
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

	classBucket := fmt.Sprintf("%s-%v", conf.GetString("GCS_BUCKET_CLASS_PREFIX"), reqClass.ID)
	if err := createBucket(classBucket); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (class ClassController) AddStudent(c *gin.Context) {
	db := database.GetDatabase()

	var studentEmails []string
	if err := c.BindJSON(&studentEmails); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	var students []models.User
	for _, studentEmail := range studentEmails {
		student, err := getUserByUserEmail(studentEmail)
		if err != nil {
			ResponseHandler(c, http.StatusInternalServerError, err.Error())
			return
		}
		students = append(students, *student)
	}

	var matchedClass models.Class
	if err := db.First(&matchedClass, c.Param("id")).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Model(&matchedClass).Association("Students").Append(&students); err != nil {
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

func (class ClassController) GetAllClasses(c *gin.Context) {
	db := database.GetDatabase()

	var allClasses []models.Class
	if err := db.Find(&allClasses).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, allClasses)
}

func (class ClassController) GetClass(c *gin.Context) {
	db := database.GetDatabase()

	var matchedClass models.Class
	if err := db.Preload("Teachers").Preload("Students").First(&matchedClass, c.Param("id")).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, matchedClass)
}

func (class ClassController) DeleteClass(c *gin.Context) {
	db := database.GetDatabase()

	var matchedClass models.Class
	if err := db.First(&matchedClass, c.Param("id")).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Delete(&matchedClass).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (class ClassController) SendInvitation(c *gin.Context) {
	envConfig := config.GetConfig()

	recipients, err := getRecipientsWithInvitationIndex(c)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	smtpSender := envConfig.GetString("SMTP_EMAIL_USERNAME")
	smtpPassword := envConfig.GetString("SMTP_EMAIL_PASSWORD")
	smtpHostname := envConfig.GetString("SMTP_HOSTNAME")
	smtpPort := envConfig.GetString("SMTP_PORT")
	smtpAddress := fmt.Sprintf("%s:%s", smtpHostname, smtpPort)
	auth := smtp.PlainAuth("", smtpSender, smtpPassword, smtpHostname)

	for invitationIndex, recipientEmail := range recipients {
		recipientHeader := fmt.Sprintf("To: %s\r\n", recipientEmail)
		subjectHeader := "Subject: Invitation to a new class!\r\n"
		// Confirmation link should  have format: https://uetclass-dev.duckdns.org/class/accept-invitation/?classId=xxx&invitationId=yyy
		body := fmt.Sprintf("Confirmation link: https://%s/class/accept-invitation/?classId=%s&invitationId=%s\r\n", envConfig.GetString("UC_DOMAIN_NAME"), c.Param("id"), invitationIndex)

		message := []byte(recipientHeader + subjectHeader + "\r\n" + body)
		if err := smtp.SendMail(smtpAddress, auth, smtpSender, []string{recipientEmail}, message); err != nil {
			ResponseHandler(c, http.StatusInternalServerError, err)
			return
		}
		fmt.Println("Email is sent to: ", recipientEmail)
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (class ClassController) AcceptInvitation(c *gin.Context) {
	db := database.GetDatabase()

	invitedClass, err := getClassByClassId(c.Query("classId"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	userEmail, err := getUserEmailByInvitationId(c.Query("invitationId"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	invitedStudent, err := getUserByUserEmail(userEmail)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Model(invitedClass).Association("Students").Append(invitedStudent); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusInternalServerError, "Succeed")
}

func (class ClassController) UploadMaterial(c *gin.Context) {
	conf := config.GetConfig()

	uploadedFile, err := c.FormFile("file")
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	bucketName := fmt.Sprintf("%s-%s", conf.GetString("GCS_BUCKET_CLASS_PREFIX"), c.Param("id"))
	if err := uploadObject(bucketName, *uploadedFile); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (class ClassController) ListMaterials(c *gin.Context) {
	conf := config.GetConfig()

	bucketName := fmt.Sprintf("%s-%s", conf.GetString("GCS_BUCKET_CLASS_PREFIX"), c.Param("id"))

	materials, err := listObjects(bucketName)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	result := map[string][]string{
		"files": materials,
	}
	ResponseHandler(c, http.StatusOK, result)
}
