package repository

import (
	"context"
	"hot_news_2/model/domain"

	"gorm.io/gorm"
)

type CommentRepository interface {
	FindByUser(ctx context.Context, tx *gorm.DB, useId int) ([]domain.Comment, error)
	FindById(ctx context.Context, tx *gorm.DB, commentId int) (domain.Comment, error)
	Create(ctx context.Context, tx *gorm.DB, comment domain.Comment) (domain.Comment, error)
	Update(ctx context.Context, tx *gorm.DB, comment domain.Comment) (domain.Comment, error)
	Delete(ctx context.Context, tx *gorm.DB, comment domain.Comment) error
}