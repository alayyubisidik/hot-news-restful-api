package service

import (
	"context"
	"hot_news_2/model/web"
)

type CategoryService interface {
	FindAll(ctx context.Context) []web.CategoryResponse
	FindBySlug(ctx context.Context, categorySlug string) web.CategoryResponse
	Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest, categorySlug string) web.CategoryResponse
	Delete(ctx context.Context, categorySlug string)
} 