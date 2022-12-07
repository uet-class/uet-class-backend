package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
)

type AssignmentController struct{}

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
