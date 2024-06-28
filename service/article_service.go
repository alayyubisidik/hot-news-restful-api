package service

import (
	"context"
	"hot_news_2/model/web"
)

type ArticleService interface {
	FindAll(ctx context.Context) []web.ArticleResponse
	FindByCategory(ctx context.Context, categorySlug string) []web.ArticleResponse
	FindByUser(ctx context.Context, username string) []web.ArticleResponse
	FindBySlug(ctx context.Context, articleSlug string) web.ArticleResponse
	Create(ctx context.Context, request web.ArticleCreateRequest) web.ArticleResponse
	Update(ctx context.Context, request web.ArticleUpdateRequest, articleSlug string) web.ArticleResponse
	Delete(ctx context.Context, articleSlug string)
} 