package business

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"github.com/liuzhaomax/maxblog-main/src/api_article/model"
	"github.com/redis/go-redis/v9"
)

var BusinessArticleSet = wire.NewSet(wire.Struct(new(BusinessArticle), "*"))

type BusinessArticle struct {
	Model *model.ModelArticle
	Tx    *core.Trans
	Redis *redis.Client
}
