package models

import (
	"time"
)

type Report struct {
	ID             int `gorm:"primaryKey"`
	ReportTypeID   int
	reportObjectID int
	CreatedAt      time.Time
	Message        string
	ReporterID     int
}
