package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzhaomax/maxblog-main/internal/middleware"
	"github.com/liuzhaomax/maxblog-main/src/api_article/handler"
)

func Register(root *gin.RouterGroup, handler *handler.HandlerArticle, mw *middleware.Middleware) {
}
