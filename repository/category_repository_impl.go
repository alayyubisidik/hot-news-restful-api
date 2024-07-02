package repository

import (
	"context"
	"hot_news_2/model/domain"

	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) ([]domain.Category, error) {
    var categories []domain.Category
    if err := tx.WithContext(ctx).Find(&categories).Error; err != nil {
        return nil, err
    }

    return categories, nil
}

func (repository *CategoryRepositoryImpl) FindBySlug(ctx context.Context, tx *gorm.DB, categorySlug string) (domain.Category, error) {
	category := domain.Category{}
	if err := tx.WithContext(ctx).Take(&category, "slug = ?", categorySlug).Error; err != nil {
		return domain.Category{}, err
	}

	return category, nil
}
 
func (repository *CategoryRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, category domain.Category) (domain.Category, error) {
	if err := tx.WithContext(ctx).Save(&category).Error; err != nil {
		return domain.Category{}, err
	}

	return category, nil
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, category domain.Category) (domain.Category, error) {
	if err := tx.WithContext(ctx).Model(&category).Updates(category).Error; err != nil {
		return domain.Category{}, err
	}

	return category, nil
} 

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, category domain.Category) error {
	if err := tx.WithContext(ctx).Delete(&category).Error; err != nil {
		return err
	}

	return nil
}
