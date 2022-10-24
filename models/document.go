package models

import (
	"gorm.io/gorm"
)

type Document struct {
	gorm.Model
	CreatorID  int
	Attachment []Attachment	`gorm:"foreignKey:DocumentID"`
	ClassID    int
	Title      string
}
