package models

import (
	"gorm.io/gorm"
)

// A Report has one reportType, ReportTypeID is the foreign key
type Report struct {
	gorm.Model
	ReportObjectID      int
	ReportObjectName    string
	ReportObjectContact string
	ReporterID          int
	ReporterName        string
	ReporterEmail       string
	ReportType          string
	Message             string
}
