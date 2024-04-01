package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/src/api_article/business"
)

var BusinessSet = wire.NewSet(
	business.BusinessArticleSet,
)
