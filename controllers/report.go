package controllers

import (
	"net/http"
	"strconv"

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

	reqUser, err := getUserByUserId(strconv.Itoa(reqReport.ReporterID))
	if err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if reqReport.ReportType == "user" {
		reportedUser, err := getUserByUserId(strconv.Itoa(reqReport.ReportObjectID))
		if err != nil {
			ResponseHandler(c, http.StatusInternalServerError, err.Error())
			return
		}
		reqReport.ReportObjectName = reportedUser.FullName
	}

	reqReport.ReporterName = reqUser.FullName
	if err := db.Create(&reqReport).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}

func (report ReportController) GetUserReports(c *gin.Context) {
	db := database.GetDatabase()
	var userReportList []models.Report

	if err := db.Where(&models.Report{ReportType: "user"}).Find(&userReportList).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, userReportList)
}

func (report ReportController) DeleteUserReport(c *gin.Context) {
	db := database.GetDatabase()

	var matchedReport models.Report
	if err := db.First(&matchedReport, c.Param("id")).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err := db.Delete(&matchedReport).Error; err != nil {
		ResponseHandler(c, http.StatusInternalServerError, err.Error())
		return
	}
	ResponseHandler(c, http.StatusOK, "Succeed")
}
