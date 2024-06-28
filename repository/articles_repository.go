package repository

import (
	"context"
	"hot_news_2/model/domain"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindAll(ctx context.Context, tx *gorm.DB) ([]domain.Article, error)
	FindByCategory(ctx context.Context, tx *gorm.DB, categoryId int) ([]domain.Article, error)
	FindByUser(ctx context.Context, tx *gorm.DB, userId int ) ([]domain.Article, error)
	FindBySlug(ctx context.Context, tx *gorm.DB, articleSlug string) (domain.Article, error)
	Create(ctx context.Context, tx *gorm.DB, article domain.Article) (domain.Article, error)
	Update(ctx context.Context, tx *gorm.DB, article domain.Article) (domain.Article, error)
	Delete(ctx context.Context, tx *gorm.DB, article domain.Article) error
}