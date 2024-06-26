package router

import (
	"github.com/gin-gonic/gin"
	"github.com/liuzhaomax/maxblog-main/internal/middleware"
	"github.com/liuzhaomax/maxblog-main/src/api_article/handler"
)

func RegisterArticle(root *gin.RouterGroup, handler *handler.HandlerArticle, mw *middleware.Middleware) {
	routerMaxBlog := root.Group("/maxblog")
	{
		routerArticle := routerMaxBlog.Group("/article")
		{
			routerArticle.GET("/list", handler.GetArticleList)
			routerArticle.GET("/tags", handler.GetArticleTags)
			routerArticle.GET("/article", handler.GetArticleByID)
			routerArticle.PATCH("/article", handler.PatchArticleLikeByID)
			routerArticle.Use(mw.Auth.ValidateToken())
			routerArticle.PUT("/article", handler.PutArticleByID)
			routerArticle.DELETE("/article", handler.DeleteArticleByID)
			routerArticle.PUT("/tag", handler.PutTagByName)
			routerArticle.DELETE("/tag", handler.DeleteTagByName)
			routerArticle.POST("/article/cover/upload", handler.PostCoverUpload)
			routerArticle.POST("/article/content/upload", handler.PostContentFileUpload)
		}
	}
}

func RegisterStatsArticle(root *gin.RouterGroup, handler *handler.HandlerStatsArticle) {
	routerMaxBlog := root.Group("/maxblog")
	{
		routerStatsArticle := routerMaxBlog.Group("/stats/article")
		{
			routerStatsArticle.GET("/main", handler.GetStatsArticleMain)
		}
	}
}
