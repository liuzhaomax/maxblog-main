package model

import (
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ArticleId string `gorm:"index:idx_article_id;unique;varchar(20);not null;"`
	Title     string `gorm:"index:idx_title;unique;varchar(150);not null;"`
	Author    string `gorm:"varchar(30);not null;"`
	Reference string `gorm:"varchar(300);"`
	Link      string `gorm:"varchar(300);"`
	View      uint   `gorm:"number;not null;default:0;"`
	Like      uint   `gorm:"number;not null;default:0;"`
	Content   string `gorm:"text;"`
	Tags      []Tag  `gorm:"many2many:article_tags;"`
}

type Tag struct {
	TagName   string `gorm:"index:idx_tag_name;varchar(20);not null;"`
	ArticleId string `gorm:"index:idx_article_id;varchar(20);"`
}
