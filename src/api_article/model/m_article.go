package model

import (
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ModelArticleSet = wire.NewSet(wire.Struct(new(ModelArticle), "*"))

type ModelArticle struct {
	DB *gorm.DB
}
