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

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *gorm.DB
	Validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *gorm.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		Validate:           validate,
	}
}

func (service *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	var categories []domain.Category

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		categories, err = service.CategoryRepository.FindAll(ctx, tx)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToCategoryResponses(categories)
}

func (service *CategoryServiceImpl) FindBySlug(ctx context.Context, categorySlug string) web.CategoryResponse {
	category := domain.Category{}

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error
		category, err = service.CategoryRepository.FindBySlug(ctx, tx, categorySlug)
		if err != nil {
			panic(exception.NewNotFoundError("Category not found"))
		}

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	slug := slug.Make(request.Name)

	category := domain.Category{
		Name: request.Name,
		Slug: slug,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		result, err := service.CategoryRepository.FindBySlug(ctx, tx, slug)
		if err == nil && result.ID != 0 {
			panic(exception.NewBadRequestError("Category is already exists"))
		}

		category, err = service.CategoryRepository.Create(ctx, tx, category)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}
 
func (service *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest, categorySlug string) web.CategoryResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	categorySlugNew := slug.Make(request.Name)

	category := domain.Category{
		Name: request.Name,
		Slug: categorySlugNew,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		existingCategory, err := service.CategoryRepository.FindBySlug(ctx, tx, categorySlug)
		if err != nil {
			panic(exception.NewNotFoundError("Category not found"))
		}

		result, err := service.CategoryRepository.FindBySlug(ctx, tx, category.Slug)
		if err == nil && result.ID != 0 && result.ID != existingCategory.ID{
			panic(exception.NewBadRequestError("Category is already exists"))
		}

		category.ID = existingCategory.ID

		category, err = service.CategoryRepository.Update(ctx, tx, category)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToCategoryResponse(category)
}

func (service *CategoryServiceImpl) Delete(ctx context.Context, categorySlug string) {
	err := service.Validate.Var(categorySlug, "required")
	helper.PanicIfError(err)

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		category, err := service.CategoryRepository.FindBySlug(ctx, tx, categorySlug)
		if err != nil {
			panic(exception.NewNotFoundError("Category not found"))
		}

		err = service.CategoryRepository.Delete(ctx, tx, category)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)
}
