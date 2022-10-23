package models

import (
	"gorm.io/gorm"
)

type Attachment struct {
	gorm.Model
	CreatorID  int
	Attachment string
}
