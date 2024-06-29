package repository

import (
	"context"
	"hot_news_2/model/domain"

	"gorm.io/gorm"
)

type LikeRepository interface {
	FindById(ctx context.Context, tx *gorm.DB, likeId int) (domain.Like, error)
	Create(ctx context.Context, tx *gorm.DB, like domain.Like) (domain.Like, error)
	Delete(ctx context.Context, tx *gorm.DB, like domain.Like) error
}