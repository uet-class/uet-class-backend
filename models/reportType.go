package models

type ReportType struct {
	ID   int `gorm:"primaryKey"`
	Name string
}
