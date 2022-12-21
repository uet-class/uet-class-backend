package controllers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
)

type SubmissionController struct{}

func (submission SubmissionController) UploadSubmission(c *gin.Context) {
	db := database.GetDatabase()

	uploadedFile, err := c.FormFile("submission")
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	assignmentId, err := strconv.Atoi(c.Query("assignmentId"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	classId, err := strconv.Atoi(c.Query("classId"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	reqUser, err := getUserByUserId(c.Query("creatorId"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	bucketName := addPrefix(os.Getenv("GCS_BUCKET_CLASS_PREFIX"), "-"+c.Query("classId"))
	assignmentPrefix := c.Query("assignmentId") + "-assignment/"
	submissionPrefix := "submissions/"
	uploadedFile.Filename = addPrefix(assignmentPrefix+submissionPrefix, uploadedFile.Filename)

	newSubmission := models.Submission{
		ClassID:      uint(classId),
		AssignmentID: uint(assignmentId),
		CreatorID:    reqUser.ID,
		CreatorName:  reqUser.FullName,
		BucketName:   bucketName,
		FileName:     uploadedFile.Filename,
	}

	if err := uploadObject(bucketName, *uploadedFile); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Create(&newSubmission).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (submission SubmissionController) GetSubmissions(c *gin.Context) {
	db := database.GetDatabase()

	var matchedSubmissions []models.Submission
	if err := db.Raw("SELECT bucket_name, file_name, creator_name FROM submissions WHERE class_id=? AND assignment_id=?", c.Query("classId"), c.Query("assignmentId")).Scan(&matchedSubmissions).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, matchedSubmissions)
}

func (submission SubmissionController) GetSubmission(c *gin.Context) {
	db := database.GetDatabase()

	var matchedSubmissions []models.Submission
	if err := db.Raw("SELECT bucket_name, file_name, creator_name FROM submissions WHERE class_id=? AND assignment_id=? AND creator_id=?", c.Query("classId"), c.Query("assignmentId"), c.Query("userId")).Scan(&matchedSubmissions).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, matchedSubmissions)
}
