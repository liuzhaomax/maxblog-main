package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"github.com/liuzhaomax/maxblog-main/src/api_article/business"
)

var HandlerArticleSet = wire.NewSet(wire.Struct(new(HandlerArticle), "*"))

type HandlerArticle struct {
	Business *business.BusinessArticle
	Logger   core.ILogger
	Res      core.IResponse
	RocketMQ core.IRocketMQ
}

func (h *HandlerArticle) GetArticleList(c *gin.Context) {

}

func (h *HandlerArticle) GetArticleTags(c *gin.Context) {

}

func (h *HandlerArticle) GetArticleByID(c *gin.Context) {

}

func (h *HandlerArticle) PutArticleByID(c *gin.Context) {

}

func (h *HandlerArticle) DeleteArticleByID(c *gin.Context) {

}

func (h *HandlerArticle) PutTagByName(c *gin.Context) {

}

func (h *HandlerArticle) DeleteTagByName(c *gin.Context) {

}
