package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CreatorID   int
	CreatorName string
	PostID      int
	Content     string
}
