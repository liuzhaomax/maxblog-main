package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/src/api_article/handler"
)

var HandlerSet = wire.NewSet(
	handler.HandlerArticleSet,
	handler.HandlerStatsArticleSet,
)
