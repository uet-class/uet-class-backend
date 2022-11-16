package models

import (
	"gorm.io/gorm"
)

// A Report has one reportType, ReportTypeID is the foreign key
type Report struct {
	gorm.Model
	ReportObjectID int
	ReporterID     int
	ReportType     string
	Message        string
}
