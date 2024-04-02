package schema

import (
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"github.com/liuzhaomax/maxblog-main/src/api_article/model"
)

type ArticleRes struct {
	Id        string   `json:"id"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
	DeletedAt string   `json:"deleted_at"`
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	Reference string   `json:"reference"`
	Link      string   `json:"link"`
	View      uint     `json:"view"`
	Like      uint     `json:"like"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
}

func MapArticle2ArticleRes(article *model.Article, tags []string) *ArticleRes {
	deletedAt := core.EmptyString
	if article.DeletedAt.Valid {
		deletedAt = article.DeletedAt.Time.String()
	}
	return &ArticleRes{
		Id:        article.ArticleId,
		CreatedAt: article.CreatedAt.String(),
		UpdatedAt: article.UpdatedAt.String(),
		DeletedAt: deletedAt,
		Title:     article.Title,
		Author:    article.Author,
		Reference: article.Reference,
		Link:      article.Link,
		View:      article.View,
		Like:      article.Like,
		Content:   article.Content,
		Tags:      tags,
	}
}
