package models

import "gorm.io/gorm"

type ReportType struct {
	gorm.Model
	ID   int
	Name string
}
