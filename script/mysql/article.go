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
