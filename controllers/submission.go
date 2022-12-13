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
		AssignmentID: uint(assignmentId),
		CreatorID:    reqUser.ID,
		CreatorName:  reqUser.FullName,
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
	bucketName := addPrefix(os.Getenv("GCS_BUCKET_CLASS_PREFIX"), "-"+c.Query("classId"))
	assignmentPrefix := c.Query("assignmentId") + "-assignment/"
	submissionPrefix := "submissions/"

	submissions, err := listObjectsWithPrefix(bucketName, assignmentPrefix+submissionPrefix)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, submissions)
}
