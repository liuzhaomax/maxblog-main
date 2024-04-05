package mysql

import (
	"github.com/lithammer/shortuuid"
	"github.com/liuzhaomax/maxblog-main/src/api_article/model"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(new(model.Article))
	if err != nil {
		return err
	}
	err = db.AutoMigrate(new(model.Tag))
	if err != nil {
		return err
	}
	return nil
}

func createArticle(db *gorm.DB) {
	for i := 0; i < 3; i++ {
		article := &model.Article{
			Model:     gorm.Model{},
			ArticleId: shortuuid.New(),
			Title:     "标题" + shortuuid.New(),
			Author:    "老马",
			Reference: "引用自我也不知道哪",
			Link:      "http://abc.com",
			View:      uint(i),
			Like:      uint(i),
			Content:   "我从小就会数数，门前大桥下，游过一群鸭，快来快来数一数，二四六七八。",
			Tags:      "abc|def|ghi|jkl",
		}
		db.Create(article)
	}
}

func createTag(db *gorm.DB) {
	tag1 := &model.Tag{
		Model:      gorm.Model{},
		Name:       "abc",
		ArticleIds: "JC23dJhf3bMNZZZCYLjGBk|BNmQim6bueJ7WQJf9pnCo4|n2aZ9oHrHDANtJw8ASVdmh",
	}
	db.Create(tag1)
	tag2 := &model.Tag{
		Model:      gorm.Model{},
		Name:       "def",
		ArticleIds: "JC23dJhf3bMNZZZCYLjGBk|BNmQim6bueJ7WQJf9pnCo4|n2aZ9oHrHDANtJw8ASVdmh",
	}
	db.Create(tag2)
	tag3 := &model.Tag{
		Model:      gorm.Model{},
		Name:       "ghi",
		ArticleIds: "JC23dJhf3bMNZZZCYLjGBk|BNmQim6bueJ7WQJf9pnCo4|n2aZ9oHrHDANtJw8ASVdmh",
	}
	db.Create(tag3)
	tag4 := &model.Tag{
		Model:      gorm.Model{},
		Name:       "jkl",
		ArticleIds: "JC23dJhf3bMNZZZCYLjGBk|BNmQim6bueJ7WQJf9pnCo4|n2aZ9oHrHDANtJw8ASVdmh",
	}
	db.Create(tag4)
}
