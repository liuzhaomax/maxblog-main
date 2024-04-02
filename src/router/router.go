package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzhaomax/maxblog-main/internal/middleware"
	"github.com/liuzhaomax/maxblog-main/src/api_article/handler"
)

func Register(root *gin.RouterGroup, handler *handler.HandlerArticle, mw *middleware.Middleware) {
	routerMaxBlog := root.Group("/maxblog")
	{
		routerArticle := routerMaxBlog.Group("/article")
		{
			routerArticle.GET("/list", handler.GetArticleList)
			routerArticle.GET("/tags", handler.GetArticleTags)
			routerArticle.GET("/:articleId", handler.GetArticleByID)
			routerArticle.Use(mw.Auth.ValidateToken())
			routerArticle.PUT("/:articleId", handler.PutArticleByID)
			routerArticle.DELETE("/:articleId", handler.DeleteArticleByID)
			routerArticle.PUT("/:tagName", handler.PutTagByName)
			routerArticle.DELETE("/:tagName", handler.DeleteTagByName)
		}
	}
}
