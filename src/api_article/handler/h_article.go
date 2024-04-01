package handler

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"github.com/liuzhaomax/maxblog-main/src/api_article/business"
)

var HandlerArticleSet = wire.NewSet(wire.Struct(new(HandlerArticle), "*"))

type HandlerArticle struct {
	Business    *business.BusinessArticle
	Logger      core.ILogger
	Res         core.IResponse
	RocketMQSet core.IRocketMQ
}
