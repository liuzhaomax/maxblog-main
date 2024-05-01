package business

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"github.com/liuzhaomax/maxblog-main/src/api_article/model"
	"github.com/liuzhaomax/maxblog-main/src/api_article/schema"
	"github.com/liuzhaomax/maxblog-main/src/utils"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"os"
	"strings"
)

var BusinessArticleSet = wire.NewSet(wire.Struct(new(BusinessArticle), "*"))

type BusinessArticle struct {
	Model *model.ModelArticle
	Tx    *core.Trans
	Redis *redis.Client
}

func (b *BusinessArticle) GetArticleList(c *gin.Context) (*[]schema.ArticleRes, error) {
	pageNoReq := c.Query(utils.PageNoQueryParamName)
	pageSizeReq := c.Query(utils.PageSizeQueryParamName)
	pageNo, pageSize, err := utils.Paginate(pageNoReq, pageSizeReq)
	if err != nil {
		return nil, core.FormatError(core.ParseIssue, "Query字段解析错误", err)
	}
	tagNamesStr := c.Query(utils.TagNameQueryParamName)
	tagNames := strings.Split(tagNamesStr, ",")
	search := c.Query(utils.SearchQueryParamName)
	list := &[]model.Article{}
	err = b.Model.QueryArticleList(c, list, pageNo, pageSize, tagNames, search)
	if err != nil {
		return nil, core.FormatError(core.DBDenied, "DB查询文章列表失败", err)
	}
	// TODO 加入缓存
	listRes := schema.MakeListRes(list)
	return listRes, nil
}

func (b *BusinessArticle) GetArticleTags(c *gin.Context) (*[]string, error) {
	tags := &[]model.Tag{}
	err := b.Model.QueryTags(c, tags)
	if err != nil {
		return nil, core.FormatError(core.DBDenied, "DB查询标签列表失败", err)
	}
	tagsRes := schema.MakeTagsRes(tags)
	return tagsRes, nil
}

func (b *BusinessArticle) GetArticleByID(c *gin.Context) (*schema.ArticleRes, error) {
	idReq := c.Query(utils.ArticleIdQueryParamName)
	article := &model.Article{}
	err := b.Tx.ExecTrans(c, func(ctx context.Context) error {
		err := b.Model.QueryArticleByID(ctx, article, idReq)
		if err != nil {
			return core.FormatError(core.DBDenied, "DB查询文章失败", err)
		}
		err = b.Model.UpdateArticleViewByID(ctx, article)
		if err != nil {
			return core.FormatError(core.DBDenied, "DB更新阅读量失败", err)
		}
		return nil
	})
	if err != nil {
		return nil, core.FormatError(core.DBDenied, "DB事务失败", err)
	}
	articleRes := schema.MapArticle2ArticleRes(article)
	return articleRes, nil
}

func (b *BusinessArticle) PatchArticleLikeByID(c *gin.Context) error {
	field := c.Query(utils.FieldQueryParamName)
	if field != "like" {
		return core.FormatError(core.MissingParameters, "参数field的值不是like", nil)
	}
	idReq := c.Query(utils.ArticleIdQueryParamName)
	article := &model.Article{}
	err := b.Tx.ExecTrans(c, func(ctx context.Context) error {
		err := b.Model.QueryArticleByID(ctx, article, idReq)
		if err != nil {
			return core.FormatError(core.DBDenied, "DB查询文章失败", err)
		}
		err = b.Model.UpdateArticleLikeByID(ctx, article)
		if err != nil {
			return core.FormatError(core.DBDenied, "DB更新点赞数失败", err)
		}
		return nil
	})
	if err != nil {
		return core.FormatError(core.DBDenied, "DB事务失败", err)
	}
	return nil
}

func (b *BusinessArticle) PutArticleByID(c *gin.Context) error {
	idReq := c.Query(utils.ArticleIdQueryParamName)
	articleReq := &schema.ArticleReq{}
	err := c.ShouldBind(articleReq)
	if err != nil {
		return core.FormatError(core.ParseIssue, "请求体无效", err)
	}
	if idReq != articleReq.Id {
		return core.FormatError(core.Forbidden, "Query信息与请求体信息不符", err)
	}
	err = validateArticleFields(articleReq)
	if err != nil {
		return core.FormatError(core.Forbidden, "校验失败", err)
	}
	article := &model.Article{}
	err = b.Tx.ExecTrans(c, func(ctx context.Context) error {
		err := b.Model.QueryArticleByID(ctx, article, articleReq.Id)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return core.FormatError(core.DBDenied, "DB查询文章失败", err)
		}
		articleToDB := schema.MapArticleReq2Article(articleReq)
		articleToDB.Model = article.Model
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = b.Model.CreateArticle(ctx, articleToDB)
			if err != nil {
				return core.FormatError(core.DBDenied, "DB创建文章失败", err)
			}
		} else {
			err = b.Model.UpdateArticleByID(ctx, articleToDB, articleReq.Id)
			if err != nil {
				return core.FormatError(core.DBDenied, "DB更新文章失败", err)
			}
		}
		return nil
	})
	if err != nil {
		return core.FormatError(core.DBDenied, "DB事务失败", err)
	}
	return nil
}

func (b *BusinessArticle) DeleteArticleByID(c *gin.Context) error {
	idReq := c.Query(utils.ArticleIdQueryParamName)
	article := &model.Article{}
	err := b.Tx.ExecTrans(c, func(ctx context.Context) error {
		err := b.Model.QueryArticleByID(ctx, article, idReq)
		if err != nil {
			return core.FormatError(core.DBDenied, "DB查询文章失败", err)
		}
		err = b.Model.DeleteArticleByID(ctx, idReq)
		if err != nil {
			return core.FormatError(core.DBDenied, "DB删除文章失败", err)
		}
		return nil
	})
	if err != nil {
		return core.FormatError(core.DBDenied, "DB事务失败", err)
	}
	return nil
}

func (b *BusinessArticle) PutTagByName(c *gin.Context) error {
	tagNameReq := c.Query(utils.TagNameQueryParamName)
	tag := &model.Tag{
		Name: tagNameReq,
	}
	err := b.Tx.ExecTrans(c, func(ctx context.Context) error {
		err := b.Model.QueryTagByName(ctx, tag, tagNameReq)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return core.FormatError(core.DBDenied, "DB查询标签失败", err)
		}
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = b.Model.CreateTag(ctx, tag)
			if err != nil {
				return core.FormatError(core.DBDenied, "DB创建标签失败", err)
			}
		} else {
			err = b.Model.UpdateTagByName(ctx, tag, tagNameReq)
			if err != nil {
				return core.FormatError(core.DBDenied, "DB更新标签失败", err)
			}
		}
		return nil
	})
	if err != nil {
		return core.FormatError(core.DBDenied, "DB事务失败", err)
	}
	return nil
}

func (b *BusinessArticle) DeleteTagByName(c *gin.Context) error {
	tagNameReq := c.Query(utils.TagNameQueryParamName)
	tag := &model.Tag{
		Name: tagNameReq,
	}
	err := b.Tx.ExecTrans(c, func(ctx context.Context) error {
		err := b.Model.QueryTagByName(ctx, tag, tagNameReq)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return core.FormatError(core.DBDenied, "DB查询标签失败", err)
		}
		err = b.Model.DeleteTagByName(ctx, tagNameReq)
		if err != nil {
			return core.FormatError(core.DBDenied, "DB删除标签失败", err)
		}
		return nil
	})
	if err != nil {
		return core.FormatError(core.DBDenied, "DB事务失败", err)
	}
	return nil
}

func (b *BusinessArticle) PostFileUpload(c *gin.Context) (string, error) {
	idReq := c.Query(utils.ArticleIdQueryParamName)
	file, err := c.FormFile(utils.File)
	if err != nil {
		return core.EmptyString, core.FormatError(core.ParseIssue, "文件读取失败", err)
	}
	dir := fmt.Sprintf("%s/%s/", utils.ArticleFileUploadPath, idReq)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return core.EmptyString, core.FormatError(core.InternalServerError, "文件路径创建失败", err)
	}
	filePath := fmt.Sprintf("%s/%s/%s", utils.ArticleFileUploadPath, idReq, file.Filename)
	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		return core.EmptyString, core.FormatError(core.InternalServerError, "文件保存失败", err)
	}
	path := fmt.Sprintf("/%s/%s/%s", utils.ArticleFileUploadPath, idReq, file.Filename)
	return path, nil
}
