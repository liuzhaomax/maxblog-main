package mysql

import (
	"github.com/liuzhaomax/maxblog-main/src/api_article/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"testing"
)

func TestAutoMigrate(t *testing.T) {
	DSN := "root:123456@tcp(127.0.0.1:3306)/maxblog?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(DSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	err := db.AutoMigrate(&model.Article{}, &model.Tag{})
	if err != nil {
		panic(err)
	}
	createArticle(db)
	createAgain(db)
}
