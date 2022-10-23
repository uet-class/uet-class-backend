package models

type Admin struct {
	ID       int `gorm:"primaryKey"`
	Email    string
	Password string
}
