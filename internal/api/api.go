package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"github.com/liuzhaomax/maxblog-main/internal/middleware"
	"github.com/liuzhaomax/maxblog-main/internal/middleware/cors"
	"github.com/liuzhaomax/maxblog-main/src/api_article/handler"
	"github.com/liuzhaomax/maxblog-main/src/router"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var APISet = wire.NewSet(wire.Struct(new(Handler), "*"), wire.Bind(new(API), new(*Handler)))

type API interface {
	Register(app *gin.Engine)
}

type Handler struct {
	Middleware          *middleware.Middleware
	HandlerArticle      *handler.HandlerArticle
	HandlerStatsArticle *handler.HandlerStatsArticle
	PrometheusRegistry  *prometheus.Registry
}

func (h *Handler) Register(app *gin.Engine) {
	// 404
	app.NoRoute(h.GetNoRoute)
	// CORS
	app.Use(cors.Cors())
	// consul
	app.GET("/health", h.HealthHandler)
	// prometheus
	app.GET("/metrics", h.MetricsHandler)
	// jaeger
	app.Use(h.Middleware.Tracing.Trace())
	// 日志
	app.Use(core.LoggerForHTTP())
	// root route
	root := app.Group("")
	{
		// interceptor
		root.Use(h.Middleware.Validator.ValidateHeaders())
		root.Use(h.Middleware.Auth.ValidateSignature())
		// dynamic api
		router.RegisterArticle(root, h.HandlerArticle, h.Middleware)
		router.RegisterStatsArticle(root, h.HandlerStatsArticle)
	}
}

func (h *Handler) RegisterStaticFS(app *gin.Engine, path string) {
	app.StaticFS("/"+path, http.Dir("./"+path))
}

func (h *Handler) GetNoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"res": "404"})
}

func (h *Handler) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"health": "ok"})
}

func (h *Handler) MetricsHandler(c *gin.Context) {
	promhttp.HandlerFor(h.PrometheusRegistry, promhttp.HandlerOpts{
		Registry: h.PrometheusRegistry,
	}).ServeHTTP(c.Writer, c.Request)
}
