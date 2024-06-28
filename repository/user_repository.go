package repository

import (
	"context"
	"hot_news_2/model/domain"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error)
	FindByUsername(ctx context.Context, tx *gorm.DB, username string) (domain.User, error)
	FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.User, error)
	Create(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
	Update(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error)
}