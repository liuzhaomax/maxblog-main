package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Id        string `gorm:"primarykey;index:idx_id;unique;varchar(20);not null;"`
	Title     string `gorm:"index:idx_title;unique;varchar(150);not null;"`
	Author    string `gorm:"varchar(30);not null;default:'马克西'"`
	Reference string `gorm:"varchar(300);"`
	Link      string `gorm:"varchar(300);"`
	View      uint   `gorm:"number;not null;default:0;"`
	Like      uint   `gorm:"number;not null;default:0;"`
	Cover     string `gorm:"varchar(300);"`
	Content   string `gorm:"text;"`
	Tags      []Tag  `gorm:"many2many:article_tag;"`
}

type Tag struct {
	Name     string    `gorm:"primarykey;index:idx_name;unique;varchar(20);not null;"`
	Articles []Article `gorm:"many2many:article_tag;"`
}
