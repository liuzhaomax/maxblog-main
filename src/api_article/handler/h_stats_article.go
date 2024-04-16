package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/internal/core"
)

var HandlerStatsArticleSet = wire.NewSet(wire.Struct(new(HandlerStatsArticle), "*"))

type HandlerStatsArticle struct {
	Logger   core.ILogger
	Res      core.IResponse
	RocketMQ core.IRocketMQ
}

func (h *HandlerStatsArticle) GetStatsArticleMain(c *gin.Context) {
	h.Res.ResSuccess(c, "ok")
}
