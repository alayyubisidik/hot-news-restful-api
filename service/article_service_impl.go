package service

import (
	"context"
	"hot_news_2/exception"
	"hot_news_2/helper"
	"hot_news_2/model/domain"
	"hot_news_2/model/web"
	"hot_news_2/repository"

	"github.com/go-playground/validator"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type ArticleServiceImpl struct {
	ArticleRepository repository.ArticleRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewArticleService(articleRepository repository.ArticleRepository, db *gorm.DB, validate *validator.Validate) ArticleService {
	return &ArticleServiceImpl{
		ArticleRepository: articleRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *ArticleServiceImpl) FindAll(ctx context.Context) []web.ArticleResponse {
	var articles []domain.Article

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		articles, err = service.ArticleRepository.FindAll(ctx, tx)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToArticleResponses(articles)
}

func (service *ArticleServiceImpl) FindByCategory(ctx context.Context, categorySlug string) []web.ArticleResponse {
	var articles []domain.Article
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		var category domain.Category
		if err := tx.WithContext(ctx).Where("slug = ?", categorySlug).First(&category).Error; err != nil {
			panic(exception.NewNotFoundError("Category not found"))
		}

		articles, err = service.ArticleRepository.FindByCategory(ctx, tx, category.ID)
		if err != nil {
			panic(exception.NewNotFoundError("Article not found"))
		}

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToArticleResponses(articles)
}

func (service *ArticleServiceImpl) FindByUser(ctx context.Context, username string) []web.ArticleResponse {
	var articles []domain.Article
	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		
		var user domain.User
		if err := tx.WithContext(ctx).Where("username = ?", username).First(&user).Error; err != nil {
			panic(exception.NewNotFoundError("User not found"))
		}

		articles, err = service.ArticleRepository.FindByUser(ctx, tx, user.ID)
		if err != nil {
			panic(exception.NewNotFoundError("Article not found"))
		}

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToArticleResponses(articles)
}

func (service *ArticleServiceImpl) FindBySlug(ctx context.Context, articleSlug string) web.ArticleResponse {
	article := domain.Article{}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		article, err = service.ArticleRepository.FindBySlug(ctx, tx, articleSlug)
		if err != nil {
			panic(exception.NewNotFoundError("Article not found"))
		}

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToArticleResponse(article)
}

func (service *ArticleServiceImpl) Create(ctx context.Context, request web.ArticleCreateRequest) web.ArticleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	slug := slug.Make(request.Title)

	article := domain.Article{
		Title: request.Title,
		Slug: slug,
		Content: request.Content,
		UserId: request.UserId,
		CategoryId: request.CategoryId,
	}  

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		var user domain.User
		if err := tx.WithContext(ctx).First(&user, article.UserId).Error; err != nil {
			panic(exception.NewNotFoundError("User not found"))
		}

		var category domain.Category
		if err := tx.WithContext(ctx).First(&category, article.CategoryId).Error; err != nil {
			panic(exception.NewNotFoundError("Category not found"))
		}

		result, err := service.ArticleRepository.FindBySlug(ctx, tx, slug)
		if err == nil && result.ID != 0 {
			panic(exception.NewBadRequestError("Article is already exists"))
		}

		article, err = service.ArticleRepository.Create(ctx, tx, article)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToArticleResponse(article)
}

func (service *ArticleServiceImpl) Update(ctx context.Context, request web.ArticleUpdateRequest, articleSlug string) web.ArticleResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	articleSlugNew := slug.Make(request.Title)

	article := domain.Article{
		Title: request.Title,
		Slug: articleSlugNew,
		Content: request.Content,
		CategoryId: request.CategoryId,
	}  

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		existingArticle, err := service.ArticleRepository.FindBySlug(ctx, tx, articleSlug)
		if err != nil {
			panic(exception.NewNotFoundError("Article not found"))
		}

		result, err := service.ArticleRepository.FindBySlug(ctx, tx, articleSlugNew)
		if err == nil && result.ID != 0 {
			panic(exception.NewBadRequestError("Article name is already exists"))
		}

		article.ID = existingArticle.ID

		article, err = service.ArticleRepository.Update(ctx, tx, article)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToArticleResponse(article)
}

func (service *ArticleServiceImpl) Delete(ctx context.Context, articleSlug string) {
	err := service.Validate.Var(articleSlug, "required")
	helper.PanicIfError(err)

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		category, err := service.ArticleRepository.FindBySlug(ctx, tx, articleSlug)
		if err != nil {
			panic(exception.NewNotFoundError("Article not found"))
		}

		err = service.ArticleRepository.Delete(ctx, tx, category)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)
}

