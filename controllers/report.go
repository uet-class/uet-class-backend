package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/uet-class/uet-class-backend/database"
	"github.com/uet-class/uet-class-backend/models"
)

type ReportController struct{}

func (report ReportController) CreateReport(c *gin.Context) {
	db := database.GetDatabase()

	var reqReport models.Report
	if err := c.BindJSON(&reqReport); err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

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

	reqReport.ReporterID = int(reqUser.ID)
	if err := db.Create(&reqReport).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (report ReportController) GetUserReports(c *gin.Context) {
	db := database.GetDatabase()
	var userReportList []models.Report

	if err := db.Where(&models.Report{ReportType: "user"}).Find(&userReportList).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err)
		return
	}
	ResponseHandler(c, http.StatusOK, userReportList)
}
