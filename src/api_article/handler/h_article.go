package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"github.com/liuzhaomax/maxblog-main/src/api_article/business"
	"net/http"
)

var HandlerArticleSet = wire.NewSet(wire.Struct(new(HandlerArticle), "*"))

type HandlerArticle struct {
	Business *business.BusinessArticle
	Logger   core.ILogger
	Res      core.IResponse
	RocketMQ core.IRocketMQ
}

func (h *HandlerArticle) GetArticleList(c *gin.Context) {
	list, err := h.Business.GetArticleList(c)
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "获取文章列表失败", err)
		return
	}
	h.Res.ResSuccess(c, *list)
}

func (h *HandlerArticle) GetArticleTags(c *gin.Context) {
	tags, err := h.Business.GetArticleTags(c)
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "获取标签列表失败", err)
		return
	}
	h.Res.ResSuccess(c, *tags)
}

func (h *HandlerArticle) GetArticleByID(c *gin.Context) {
	article, err := h.Business.GetArticleByID(c)
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "获取文章失败", err)
		return
	}
	h.Res.ResSuccess(c, *article)
}

func (h *HandlerArticle) PatchArticleLikeByID(c *gin.Context) {
	err := h.Business.PatchArticleLikeByID(c)
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "获取文章失败", err)
		return
	}
	h.Res.ResSuccess(c, "ok")
}

func (h *HandlerArticle) PutArticleByID(c *gin.Context) {
	err := h.Business.PutArticleByID(c)
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "更新文章失败", err)
		return
	}
	h.Res.ResSuccess(c, "ok")
}

func (h *HandlerArticle) DeleteArticleByID(c *gin.Context) {
	err := h.Business.DeleteArticleByID(c)
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "删除文章失败", err)
		return
	}
	h.Res.ResSuccess(c, "ok")
}

func (h *HandlerArticle) PutTagByName(c *gin.Context) {
	err := h.Business.PutTagByName(c)
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "更新标签失败", err)
		return
	}
	h.Res.ResSuccess(c, "ok")

}

func (h *HandlerArticle) DeleteTagByName(c *gin.Context) {
	err := h.Business.DeleteTagByName(c)
	if err != nil {
		h.Res.ResFailure(c, http.StatusInternalServerError, core.InternalServerError, "删除标签失败", err)
		return
	}
	h.Res.ResSuccess(c, "ok")
}
