package models

import (
	"gorm.io/gorm"
	"time"
)

type Report struct {
	gorm.Model
	ID             int
	ReportTypeID   int
	reportObjectID int
	CreatedAt      time.Time
	Message        string
	ReporterID     int
}
