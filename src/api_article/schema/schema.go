package schema

import (
	"github.com/liuzhaomax/maxblog-main/internal/core"
	"github.com/liuzhaomax/maxblog-main/src/api_article/model"
	"github.com/liuzhaomax/maxblog-main/src/api_article/pb"
)

type ArticleReq struct {
	Id        string   `json:"id"`
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	Reference string   `json:"reference"`
	Link      string   `json:"link"`
	View      uint     `json:"view"`
	Like      uint     `json:"like"`
	Cover     string   `json:"cover"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
}

func MapArticleReq2Article(articleReq *ArticleReq) *model.Article {
	article := &model.Article{
		Id:        articleReq.Id,
		Title:     articleReq.Title,
		Author:    articleReq.Author,
		Reference: articleReq.Reference,
		Link:      articleReq.Link,
		View:      articleReq.View,
		Like:      articleReq.Like,
		Cover:     articleReq.Cover,
		Content:   articleReq.Content,
		Tags:      []model.Tag{},
	}
	for _, tag := range articleReq.Tags {
		article.Tags = append(article.Tags, model.Tag{
			Name: tag,
		})
	}
	if article.Id == core.EmptyString {
		article.Id = core.ShortUUID()
	}
	return article
}

type ArticleRes struct {
	Id        string   `json:"id"`
	CreatedAt string   `json:"createdAt"`
	UpdatedAt string   `json:"updatedAt"`
	DeletedAt string   `json:"deletedAt"`
	Title     string   `json:"title"`
	Author    string   `json:"author"`
	Reference string   `json:"reference"`
	Link      string   `json:"link"`
	View      uint     `json:"view"`
	Like      uint     `json:"like"`
	Cover     string   `json:"cover"`
	Content   string   `json:"content"`
	Tags      []string `json:"tags"`
}

func MapArticle2ArticleRes(article *model.Article) *ArticleRes {
	deletedAt := core.EmptyString
	if article.DeletedAt.Valid {
		deletedAt = article.DeletedAt.Time.String()
	}
	tagNames := []string{}
	for _, tag := range article.Tags {
		tagNames = append(tagNames, tag.Name)
	}
	return &ArticleRes{
		Id:        article.Id,
		CreatedAt: article.CreatedAt.String(),
		UpdatedAt: article.UpdatedAt.String(),
		DeletedAt: deletedAt,
		Title:     article.Title,
		Author:    article.Author,
		Reference: article.Reference,
		Link:      article.Link,
		View:      article.View,
		Like:      article.Like,
		Cover:     article.Cover,
		Content:   article.Content,
		Tags:      tagNames,
	}
}

func MakeListRes(list *[]model.Article) *[]ArticleRes {
	listRes := []ArticleRes{}
	for _, article := range *list {
		listRes = append(listRes, *MapArticle2ArticleRes(&article))
	}
	return &listRes
}

// Tag

func MapTag2TagRes(tag *model.Tag) string {
	return tag.Name
}

func MakeTagsRes(tags *[]model.Tag) *[]string {
	tagsRes := []string{}
	for _, tag := range *tags {
		tagsRes = append(tagsRes, MapTag2TagRes(&tag))
	}
	return &tagsRes
}

// Stats

type StatsArticleMainRes struct {
	Quantity int `json:"quantity"`
	View     int `json:"view"`
	Like     int `json:"like"`
}

func MapStatsArticleMainRes(rpcRes *pb.StatsArticleMainRes) *StatsArticleMainRes {
	return &StatsArticleMainRes{
		Quantity: int(rpcRes.Quantity),
		View:     int(rpcRes.View),
		Like:     int(rpcRes.Like),
	}
}
