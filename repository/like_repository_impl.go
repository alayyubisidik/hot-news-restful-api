package repository

import (
	"context"
	"hot_news_2/model/domain"

	"gorm.io/gorm"
)

type LikeRepositoryImpl struct {
}

func NewLikeRepository() LikeRepository {
	return &LikeRepositoryImpl{}
}

func (repository *LikeRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, likeId int) (domain.Like, error) {
	like := domain.Like{}
	if err := tx.WithContext(ctx).Take(&like, "id = ?", likeId).Error; err != nil {
		return domain.Like{}, err
	}

	return like, nil
}

func (repository *LikeRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, like domain.Like) (domain.Like, error) {
	if err := tx.WithContext(ctx).Save(&like).Error; err != nil {
		return domain.Like{}, err
	}

	if err := tx.WithContext(ctx).Preload("User").Preload("Article").Take(&like, "id = ?", like.ID).Error; err != nil {
		return domain.Like{}, err
	}

	return like, nil
}

func (repository *LikeRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, like domain.Like) error {
	if err := tx.WithContext(ctx).Delete(&like).Error; err != nil {
		return err
	}

	return nil
}
