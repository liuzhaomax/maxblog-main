package schema

import (
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"github.com/liuzhaomax/maxblog-main/src/api_article/model"
)

type ArticleReq struct {
	ArticleId string `json:"articleId"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Reference string `json:"reference"`
	Link      string `json:"link"`
	View      uint   `json:"view"`
	Like      uint   `json:"like"`
	Content   string `json:"content"`
	Tags      string `json:"tags"`
}

func MapArticleReq2Article(articleReq *ArticleReq) *model.Article {
	article := &model.Article{
		ArticleId: articleReq.ArticleId,
		Title:     articleReq.Title,
		Author:    articleReq.Author,
		Reference: articleReq.Reference,
		Link:      articleReq.Link,
		View:      articleReq.View,
		Like:      articleReq.Like,
		Content:   articleReq.Content,
		Tags:      articleReq.Tags,
	}
	return article
}

type ArticleRes struct {
	ArticleId string `json:"articleId"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	DeletedAt string `json:"deletedAt"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Reference string `json:"reference"`
	Link      string `json:"link"`
	View      uint   `json:"view"`
	Like      uint   `json:"like"`
	Content   string `json:"content"`
	Tags      string `json:"tags"`
}

func MapArticle2ArticleRes(article *model.Article) *ArticleRes {
	deletedAt := core.EmptyString
	if article.DeletedAt.Valid {
		deletedAt = article.DeletedAt.Time.String()
	}
	return &ArticleRes{
		ArticleId: article.ArticleId,
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
		Tags:      article.Tags,
	}
}

func MakeListRes(list *[]model.Article) *[]ArticleRes {
	listRes := []ArticleRes{}
	for _, article := range *list {
		listRes = append(listRes, *MapArticle2ArticleRes(&article))
	}
	return &listRes
}

type TagReq struct {
	Name       string `json:"name"`
	ArticleIds string `json:"articleIds"`
}

func MapTagReq2Tag(tagReq *TagReq) *model.Tag {
	tag := &model.Tag{
		Name:       tagReq.Name,
		ArticleIds: tagReq.ArticleIds,
	}
	return tag
}

type TagRes struct {
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
	DeletedAt  string `json:"deletedAt"`
	Name       string `json:"name"`
	ArticleIds string `json:"articleIds"`
}

func MapTag2TagRes(tag *model.Tag) *TagRes {
	deletedAt := core.EmptyString
	if tag.DeletedAt.Valid {
		deletedAt = tag.DeletedAt.Time.String()
	}
	return &TagRes{
		CreatedAt:  tag.CreatedAt.String(),
		UpdatedAt:  tag.UpdatedAt.String(),
		DeletedAt:  deletedAt,
		Name:       tag.Name,
		ArticleIds: tag.ArticleIds,
	}
}

func MakeTagsRes(tags *[]model.Tag) *[]TagRes {
	tagsRes := []TagRes{}
	for _, tag := range *tags {
		tagsRes = append(tagsRes, *MapTag2TagRes(&tag))
	}
	return &tagsRes
}
