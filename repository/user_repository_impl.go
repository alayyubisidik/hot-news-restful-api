package repository

import (
	"context"
	"hot_news_2/model/domain"

	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *gorm.DB, userId int) (domain.User, error) {
	user := domain.User{}
	if err := tx.WithContext(ctx).Take(&user, "id = ?", userId).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *gorm.DB, username string) (domain.User, error) {
	user := domain.User{}
	if err := tx.WithContext(ctx).Take(&user, "username = ?", username).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (domain.User, error) {
	user := domain.User{}
	if err := tx.WithContext(ctx).Take(&user, "email = ?", email).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error) {
	if err := tx.WithContext(ctx).Save(&user).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, user domain.User) (domain.User, error) {
	if err := tx.WithContext(ctx).Model(&user).Omit("password").Updates(user).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
} 