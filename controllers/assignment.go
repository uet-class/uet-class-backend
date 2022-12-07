package controllers

import (
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
)

type AssignmentController struct{}

func addPrefix(prefix string, objectName string) string {
	return prefix + objectName
}

func (assignment AssignmentController) CreateAssignment(c *gin.Context) {
	db := database.GetDatabase()

	var reqAssignment models.Assignment
	if err := c.BindJSON(&reqAssignment); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	matchedUser, err := getUserByUserId(strconv.Itoa(reqAssignment.CreatorID))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	reqAssignment.CreatorName = matchedUser.FullName

	if err := db.Create(&reqAssignment).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (assignment AssignmentController) GetAssignments(c *gin.Context) {
	db := database.GetDatabase()

	classId, err := strconv.Atoi(c.Query("classId"))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	var assignments []models.Assignment
	if err := db.Find(&assignments, models.Assignment{ClassID: classId}).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, assignments)
}

func (assignment AssignmentController) UploadAttachment(c *gin.Context) {
	uploadedFile, err := c.FormFile("file")
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	bucketName := addPrefix(os.Getenv("GCS_BUCKET_CLASS_PREFIX"), "-"+c.Query("classId"))
	uploadedFile.Filename = addPrefix(c.Param("id")+"-assignment/", uploadedFile.Filename)
	if err := uploadObject(bucketName, *uploadedFile); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (assignment AssignmentController) GetAssignment(c *gin.Context) {
	db := database.GetDatabase()

	var matchedAssignment models.Assignment
	if err := db.First(&matchedAssignment, c.Param("id")).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	bucketName := addPrefix(os.Getenv("GCS_BUCKET_CLASS_PREFIX"), "-"+c.Query("classId"))
	subDirPrefix := c.Param("id") + "-assignment/"

	objectList, err := listObjectsWithPrefix(bucketName, subDirPrefix)
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	result := map[string]interface{}{
		"Assignment": matchedAssignment,
		"UploadFile": objectList,
	}

	ResponseHandler(c, http.StatusOK, result)
}
