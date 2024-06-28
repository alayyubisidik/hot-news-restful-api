package helper

import (
	"hot_news_2/model/domain"
	"hot_news_2/model/web"
)

func ToUserResponse(user domain.User) web.UserResponse {
	return web.UserResponse{
		Id:        user.ID,
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt,
	}
}

func ToArticleResponse(article domain.Article) web.ArticleResponse {
	return web.ArticleResponse{
		Id:        article.ID,
		Title:     article.Title,
		Slug:      article.Slug,
		Content:   article.Content,
		CreatedAt: article.CreatedAt,
		User: web.UserResponse{
			Id:        article.User.ID,
			Username:  article.User.Username,
			FullName:  article.User.FullName,
			Email:     article.User.Email,
			CreatedAt: article.User.CreatedAt,
		},
		Category: web.CategoryResponse{
			Id:   article.Category.ID,
			Name: article.Category.Name,
			Slug: article.Category.Slug,
		},
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}
	return categoryResponses
}

func ToArticleResponses(articles []domain.Article) []web.ArticleResponse {
	var articleResponses []web.ArticleResponse
	for _, article := range articles {
		articleResponses = append(articleResponses, ToArticleResponse(article))
	}
	return articleResponses
}
