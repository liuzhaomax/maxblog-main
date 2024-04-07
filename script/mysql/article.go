package mysql

import (
	"fmt"
	"github.com/liuzhaomax/maxblog-main/src/api_article/model"
	"gorm.io/gorm"
)

func createArticle(db *gorm.DB) {
	articleIds := []string{"JC23dJhf3bMNZZZCYLjGBk", "BNmQim6bueJ7WQJf9pnCo4", "n2aZ9oHrHDANtJw8ASVdmh"}
	article0 := &model.Article{
		ArticleId: articleIds[0],
		Title:     fmt.Sprintf("标题-%d", 0),
		Author:    "老马",
		Reference: "引用自我也不知道哪",
		Link:      "http://abc.com",
		View:      uint(0),
		Like:      uint(0),
		Content:   "我从小就会数数，门前大桥下，游过一群鸭，快来快来数一数，二四六七八。",
		Tags: []model.Tag{
			{
				Name: "abc",
			},
			{
				Name: "def",
			},
		},
	}
	db.Create(article0)
	article1 := &model.Article{
		ArticleId: articleIds[1],
		Title:     fmt.Sprintf("标题-%d", 1),
		Author:    "老马",
		Reference: "引用自我也不知道哪",
		Link:      "http://abc.com",
		View:      uint(1),
		Like:      uint(1),
		Content:   "我从小就会数数，门前大桥下，游过一群鸭，快来快来数一数，二四六七八。",
		Tags: []model.Tag{
			{
				Name: "abc",
			},
		},
	}
	db.Create(article1)
	article2 := &model.Article{
		ArticleId: articleIds[2],
		Title:     fmt.Sprintf("标题-%d", 2),
		Author:    "老马",
		Reference: "引用自我也不知道哪",
		Link:      "http://abc.com",
		View:      uint(2),
		Like:      uint(2),
		Content:   "我从小就会数数，门前大桥下，游过一群鸭，快来快来数一数，二四六七八。",
		Tags:      []model.Tag{},
	}
	db.Create(article2)
}

func createAgain(db *gorm.DB) {
	article := &model.Article{
		ArticleId: "ooooooooo",
		Title:     fmt.Sprintf("标题-%d", 3),
		Author:    "老马",
		Reference: "引用自我也不知道哪",
		Link:      "http://abc.com",
		View:      uint(3),
		Like:      uint(3),
		Content:   "我从小就会数数，门前大桥下，游过一群鸭，快来快来数一数，二四六七八。",
		Tags: []model.Tag{
			{
				Name: "abc",
			},
		},
	}
	db.Create(article)
}
