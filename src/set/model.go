package set

import (
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/src/api_article/model"
)

var ModelSet = wire.NewSet(
	model.ModelArticleSet,
)
