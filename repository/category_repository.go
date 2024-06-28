package repository

import (
	"context"
	"hot_news_2/model/domain"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(ctx context.Context, tx *gorm.DB) ([]domain.Category, error)
	FindBySlug(ctx context.Context, tx *gorm.DB, categorySlug string) (domain.Category, error)
	Create(ctx context.Context, tx *gorm.DB, category domain.Category) (domain.Category, error)
	Update(ctx context.Context, tx *gorm.DB, category domain.Category) (domain.Category, error)
	Delete(ctx context.Context, tx *gorm.DB, category domain.Category) error
}