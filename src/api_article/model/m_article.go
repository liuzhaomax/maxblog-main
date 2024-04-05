package model

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var ModelArticleSet = wire.NewSet(wire.Struct(new(ModelArticle), "*"))

type ModelArticle struct {
	DB *gorm.DB
}

func (m *ModelArticle) QueryArticleList(c *gin.Context, list *[]Article, pageNo int, pageSize int) error {
	offset := (pageNo - 1) * pageSize
	scope := func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
	result := m.DB.WithContext(c).Scopes(scope).Find(list)
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

func (m *ModelArticle) QueryArticleByID(c *gin.Context, article *Article, articleId string) error {
	result := m.DB.WithContext(c).Where("article_id=?", articleId).First(article)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) CreateArticle(c *gin.Context, article *Article) error {
	result := m.DB.WithContext(c).Create(article)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) UpdateArticleByID(c *gin.Context, article *Article, articleId string) error {
	result := m.DB.WithContext(c).Where("article_id=?", articleId).Updates(article)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) DeleteArticleByID(c *gin.Context, articleId string) error {
	result := m.DB.WithContext(c).Where("article_id=?", articleId).Delete(&Article{})
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) QueryTagByName(c *gin.Context, tag *Tag, tagName string) error {
	result := m.DB.WithContext(c).Where("name=?", tagName).First(tag)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) CreateTag(c *gin.Context, tag *Tag) error {
	result := m.DB.WithContext(c).Create(tag)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) UpdateTagByName(c *gin.Context, tag *Tag, tagName string) error {
	result := m.DB.WithContext(c).Where("name=?", tagName).Updates(tag)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) DeleteTagByName(c *gin.Context, tagName string) error {
	result := m.DB.WithContext(c).Where("name=?", tagName).Delete(&Tag{})
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}
