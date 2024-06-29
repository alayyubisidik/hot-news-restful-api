package service

import (
	"context"
	"hot_news_2/exception"
	"hot_news_2/helper"
	"hot_news_2/model/domain"
	"hot_news_2/model/web"
	"hot_news_2/repository"

	"github.com/go-playground/validator"
	"gorm.io/gorm"
)

type LikeServiceImpl struct {
	LikeRepository repository.LikeRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewLikeService(LikeRepository repository.LikeRepository, db *gorm.DB, validate *validator.Validate) LikeService {
	return &LikeServiceImpl{
		LikeRepository: LikeRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *LikeServiceImpl) FindById(ctx context.Context, likeId int) web.LikeResponse {
	var like domain.Like

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		like, err = service.LikeRepository.FindById(ctx, tx, likeId)
		if err != nil {
			panic(exception.NewNotFoundError("Like not found"))
		}

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToLikeResponse(like)
}

func (service *LikeServiceImpl) Create(ctx context.Context, request web.LikeCreateRequest) web.LikeResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	like := domain.Like{
		UserId:    request.UserId,
		ArticleId: request.ArticleId,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		var user domain.User
		if err := tx.WithContext(ctx).First(&user, like.UserId).Error; err != nil {
			panic(exception.NewNotFoundError("User not found"))
		}

		var article domain.Article
		if err := tx.WithContext(ctx).First(&article, like.ArticleId).Error; err != nil {
			panic(exception.NewNotFoundError("Article not found"))
		}

		like, err = service.LikeRepository.Create(ctx, tx, like)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToLikeResponse(like)
}

func (service *LikeServiceImpl) Delete(ctx context.Context, likeId int) {
	err := service.Validate.Var(likeId, "required")
	helper.PanicIfError(err)

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		like, err := service.LikeRepository.FindById(ctx, tx, likeId)
		if err != nil {
			panic(exception.NewNotFoundError("Like not found"))
		}

		err = service.LikeRepository.Delete(ctx, tx, like)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)
}

