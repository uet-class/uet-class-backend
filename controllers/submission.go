package controllers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type SubmissionController struct{}

func (submission SubmissionController) UploadSubmission(c *gin.Context) {
	uploadedFile, err := c.FormFile("submission")
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	bucketName := addPrefix(os.Getenv("GCS_BUCKET_CLASS_PREFIX"), "-"+c.Query("classId"))
	assignmentPrefix := c.Query("assignmentId") + "-assignment/"
	submissionPrefix := "submissions/"
	uploadedFile.Filename = addPrefix(assignmentPrefix+submissionPrefix, uploadedFile.Filename)
	
	if err := uploadObject(bucketName, *uploadedFile); err != nil {
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
