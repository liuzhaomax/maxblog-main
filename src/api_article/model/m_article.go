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

func (m *ModelArticle) QueryArticleList(c *gin.Context, list *[]Article, pageNo int, pageSize int, tagNames []string) error {
	offset := (pageNo - 1) * pageSize
	scope := func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
	query := m.DB.WithContext(c).Scopes(scope)
	if len(tagNames) > 0 {
		query = query.
			Select("DISTINCT article.*").
			Joins("JOIN article_tag ON article.id = article_tag.article_id").
			Joins("JOIN tag ON article_tag.tag_name = tag.name").
			Where("tag.name IN ?", tagNames)
	}
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

func (m *ModelArticle) QueryArticleByID(c *gin.Context, article *Article, articleId string) error {
	result := m.DB.WithContext(c).Preload("Tags").Where("id=?", articleId).First(article)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) CreateArticle(c *gin.Context, article *Article) error {
	result := m.DB.WithContext(c).Preload("Tags").Create(article)
	if result.RowsAffected == 0 {
		return result.Error
	}
	return nil
}

func (m *ModelArticle) UpdateArticleByID(c *gin.Context, article *Article, articleId string) error {
	tx := m.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	result := tx.WithContext(c).Preload("Tags").Where("id=?", articleId).Updates(article)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	err := tx.WithContext(c).Model(article).Association("Tags").Replace(&article.Tags)
	if err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (m *ModelArticle) DeleteArticleByID(c *gin.Context, articleId string) error {
	tx := m.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	article := &Article{
		Id: articleId, // Clear函数必须指定主键
	}
	err := tx.WithContext(c).Model(article).Where("article_id = ?", articleId).
		Association("Tags").Clear()
	if err != nil {
		tx.Rollback()
		return err
	}
	result := tx.WithContext(c).Where("id=?", articleId).Delete(&Article{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
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
	tx := m.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	tag := &Tag{
		Name: tagName,
	}
	err := tx.WithContext(c).Model(tag).Where("tag_name = ?", tagName).
		Association("Articles").Clear()
	if err != nil {
		tx.Rollback()
		return err
	}
	result := tx.WithContext(c).Where("name=?", tagName).Delete(&Tag{})
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}
	return tx.Commit().Error
}
