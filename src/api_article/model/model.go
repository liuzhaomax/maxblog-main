package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title string `gorm:"index:idx_title;unique;varchar(30);not null"`
}
