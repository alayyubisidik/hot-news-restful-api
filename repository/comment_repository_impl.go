package repository

import (
	"context"
	"hot_news_2/model/domain"

	"gorm.io/gorm"
)

type CommentRepositoryImpl struct {
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}

func (repository *CommentRepositoryImpl) FindByUser(ctx context.Context, tx *gorm.DB, userId int) ([]domain.Comment, error) {
	comments := []domain.Comment{}
	if err := tx.WithContext(ctx).Preload("User").Where("user_id = ?", userId).Find(&comments).Error; err != nil {
		return []domain.Comment{}, err
	}

	return comments, nil
}

func (repository *CommentRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, commentId int) (domain.Comment, error) {
	comment := domain.Comment{}
	if err := tx.WithContext(ctx).Preload("User").Preload("Article").Take(&comment, "id = ?", commentId).Error; err != nil {
		return domain.Comment{}, err
	}

	return comment, nil
}

func (repository *CommentRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, comment domain.Comment) (domain.Comment, error) {
	if err := tx.WithContext(ctx).Save(&comment).Error; err != nil {
		return domain.Comment{}, err
	}

	if err := tx.WithContext(ctx).Preload("User").Preload("Article").First(&comment, comment.ID).Error; err != nil {
		return domain.Comment{}, err
	}

	return comment, nil
}

func (repository *CommentRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, comment domain.Comment) (domain.Comment, error) {
	if err := tx.WithContext(ctx).Model(&comment).Updates(comment).Error; err != nil {
		return domain.Comment{}, err
	}

	if err := tx.WithContext(ctx).Preload("User").Preload("Article").First(&comment, comment.ID).Error; err != nil {
		return domain.Comment{}, err
	}

	return comment, nil
}

func (repository *CommentRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, comment domain.Comment) error {
	if err := tx.WithContext(ctx).Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}