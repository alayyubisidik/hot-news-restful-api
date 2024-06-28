package repository

import (
	"context"
	"hot_news_2/model/domain"

	"gorm.io/gorm"
)

type ArticleRepositoryImpl struct {
}

func NewArticleRepository() ArticleRepository {
	return &ArticleRepositoryImpl{}
}

func (repository *ArticleRepositoryImpl) FindAll(ctx context.Context, tx *gorm.DB) ([]domain.Article, error) {
	var articles []domain.Article
	if err := tx.WithContext(ctx).Preload("Category").Preload("User").Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (repository *ArticleRepositoryImpl) FindByCategory(ctx context.Context, tx *gorm.DB, categoryId int) ([]domain.Article, error) {
	articles := []domain.Article{}
	if err := tx.WithContext(ctx).Preload("Category").Preload("User").Where("category_id = ?", categoryId).Find(&articles).Error; err != nil {
		return []domain.Article{}, err
	}

	return articles, nil
}

func (repository *ArticleRepositoryImpl) FindByUser(ctx context.Context, tx *gorm.DB, userId int) ([]domain.Article, error) {
	articles := []domain.Article{}
	if err := tx.WithContext(ctx).Preload("Category").Preload("User").Where("user_id = ?", userId).Find(&articles).Error; err != nil {
		return []domain.Article{}, err
	}

	return articles, nil
}

func (repository *ArticleRepositoryImpl) FindBySlug(ctx context.Context, tx *gorm.DB, articleSlug string) (domain.Article, error) {
	article := domain.Article{}
	if err := tx.WithContext(ctx).Joins("Category").Joins("User").Take(&article, "articles.slug = ?", articleSlug).Error; err != nil {
		return domain.Article{}, err
	}

	return article, nil
}

func (repository *ArticleRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, article domain.Article) (domain.Article, error) {
	if err := tx.WithContext(ctx).Save(&article).Error; err != nil {
		return domain.Article{}, err
	}

	if err := tx.WithContext(ctx).Preload("User").Preload("Category").First(&article, article.ID).Error; err != nil {
		return domain.Article{}, err
	}

	return article, nil
}


func (repository *ArticleRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, article domain.Article) (domain.Article, error) {
	if err := tx.WithContext(ctx).Model(&article).Updates(article).Error; err != nil {
		return domain.Article{}, err
	}

	if err := tx.WithContext(ctx).Preload("User").Preload("Category").First(&article, article.ID).Error; err != nil {
		return domain.Article{}, err
	}

	return article, nil
}

func (repository *ArticleRepositoryImpl) Delete(ctx context.Context, tx *gorm.DB, article domain.Article) error {
	if err := tx.WithContext(ctx).Delete(&article).Error; err != nil {
		return err
	}

	return nil
}
