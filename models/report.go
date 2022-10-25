package models

import (
	"gorm.io/gorm"
)

// A Report has one reportType, ReportTypeID is the foreign key
type Report struct {
	gorm.Model
	ReportTypeID   int
	ReportObjectID int
	Message        string
	ReporterID     int
	ReportType     ReportType
}
