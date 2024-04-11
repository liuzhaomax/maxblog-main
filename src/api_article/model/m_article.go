package model

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"gorm.io/gorm"
	"sync/atomic"
)

var ModelArticleSet = wire.NewSet(wire.Struct(new(ModelArticle), "*"))

type ModelArticle struct {
	DB *gorm.DB
}

func (m *ModelArticle) QueryArticleList(c *gin.Context, list *[]Article, pageNo int, pageSize int, tagNames []string, search string) error {
	offset := (pageNo - 1) * pageSize
	scope := func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
	query := m.DB.WithContext(c).Scopes(scope)
	// tagNames元素不全为空字符串
	var isAllEmptyStringElem bool
	for _, tag := range tagNames {
		if tag != core.EmptyString {
			isAllEmptyStringElem = true
			break
		}
	}
	if len(tagNames) > 0 && isAllEmptyStringElem {
		query = query.
			Select("DISTINCT article.*").
			Joins("JOIN article_tag ON article.id = article_tag.article_id").
			Joins("JOIN tag ON article_tag.tag_name = tag.name").
			Where("tag.name IN ?", tagNames)
	}
	if search != core.EmptyString {
		searchCond := fmt.Sprintf("%%%s%%", search)
		query = query.Where("article.title LIKE ? OR article.content LIKE ?", searchCond, searchCond)
	}
	query = query.Order("article.updated_at DESC")
	result := query.Preload("Tags").Find(list)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) QueryTags(c *gin.Context, tags *[]Tag) error {
	result := m.DB.WithContext(c).Find(tags)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) QueryArticleByID(ctx context.Context, article *Article, articleId string) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).
		Set("gorm:query_option", "FOR UPDATE").
		Preload("Tags").
		Where("id=?", articleId).
		First(article)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) UpdateArticleViewByID(ctx context.Context, article *Article) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	view := uint32(article.View)
	atomic.AddUint32(&view, 1)
	result := tx.WithContext(ctx).Model(article).
		Set("gorm:query_option", "FOR UPDATE").
		Update("view", atomic.LoadUint32(&view))
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) CreateArticle(ctx context.Context, article *Article) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).Preload("Tags").Create(article)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) UpdateArticleByID(ctx context.Context, article *Article, articleId string) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).
		Set("gorm:query_option", "FOR UPDATE").
		Preload("Tags").
		Where("id=?", articleId).
		Updates(article)
	if result.Error != nil {
		return result.Error
	}
	err := tx.WithContext(ctx).Model(article).
		Set("gorm:query_option", "FOR UPDATE").
		Association("Tags").
		Replace(&article.Tags)
	if err != nil {
		return err
	}
	return nil
}

func (m *ModelArticle) DeleteArticleByID(ctx context.Context, articleId string) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	article := &Article{
		Id: articleId, // Clear函数必须指定主键
	}
	err := tx.WithContext(ctx).Model(article).
		Set("gorm:query_option", "FOR UPDATE").
		Where("article_id = ?", articleId).
		Association("Tags").
		Clear()
	if err != nil {
		return err
	}
	result := tx.WithContext(ctx).Where("id=?", articleId).Delete(&Article{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) QueryTagByName(ctx context.Context, tag *Tag, tagName string) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).
		Set("gorm:query_option", "FOR UPDATE").
		Where("name=?", tagName).
		First(tag)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) CreateTag(ctx context.Context, tag *Tag) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).Create(tag)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) UpdateTagByName(ctx context.Context, tag *Tag, tagName string) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	result := tx.WithContext(ctx).
		Set("gorm:query_option", "FOR UPDATE").
		Where("name=?", tagName).
		Updates(tag)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) DeleteTagByName(ctx context.Context, tagName string) error {
	tx := ctx.Value(core.Trans{}).(*gorm.DB)
	tag := &Tag{
		Name: tagName,
	}
	err := tx.WithContext(ctx).Model(tag).
		Set("gorm:query_option", "FOR UPDATE").
		Where("tag_name = ?", tagName).
		Association("Articles").
		Clear()
	if err != nil {
		return err
	}
	result := tx.WithContext(ctx).
		Set("gorm:query_option", "FOR UPDATE").
		Where("name=?", tagName).
		Delete(&Tag{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}
