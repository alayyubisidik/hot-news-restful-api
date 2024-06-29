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
 
type CommentServiceImpl struct {
	CommentRepository repository.CommentRepository
	DB                *gorm.DB
	Validate          *validator.Validate
}

func NewCommentService(commentRepository repository.CommentRepository, db *gorm.DB, validate *validator.Validate) CommentService {
	return &CommentServiceImpl{
		CommentRepository: commentRepository,
		DB: db,
		Validate: validate,
	}
}

func (service *CommentServiceImpl) FindByUser(ctx context.Context, username string) []web.CommentResponse {
	var comments []domain.Comment

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		var user domain.User
		if err = tx.WithContext(ctx).Take(&user, "username = ?", username).Error; err != nil {
			panic(exception.NewNotFoundError("User not found"))
		}

		comments, err = service.CommentRepository.FindByUser(ctx, tx, user.ID)
		if err != nil {
			panic(exception.NewNotFoundError("Comment not found"))
		}

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToCommentResponses(comments)
}

func (service *CommentServiceImpl) FindById(ctx context.Context, commentId int) web.CommentResponse {
	var comment domain.Comment

	err := service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		comment, err = service.CommentRepository.FindById(ctx, tx, commentId)
		if err != nil {
			panic(exception.NewNotFoundError("Comment not found"))
		}

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToCommentResponse(comment)
}

func (service *CommentServiceImpl) Create(ctx context.Context, request web.CommentCreateRequest) web.CommentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	comment := domain.Comment{
		UserId:    request.UserId,
		ArticleId: request.ArticleId,
		Content:   request.Content,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		var user domain.User
		if err := tx.WithContext(ctx).First(&user, comment.UserId).Error; err != nil {
			panic(exception.NewNotFoundError("User not found"))
		}

		var article domain.Article
		if err := tx.WithContext(ctx).First(&article, comment.ArticleId).Error; err != nil {
			panic(exception.NewNotFoundError("Article not found"))
		}

		comment, err = service.CommentRepository.Create(ctx, tx, comment)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToCommentResponse(comment)
}

func (service *CommentServiceImpl) Update(ctx context.Context, request web.CommentUpdateRequest) web.CommentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	comment := domain.Comment{
		ID: request.Id,
		Content:   request.Content,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		var err error

		_, err = service.CommentRepository.FindById(ctx, tx, comment.ID)
		if err != nil  {
			panic(exception.NewNotFoundError("Comment not found"))
		}

		comment, err = service.CommentRepository.Update(ctx, tx, comment)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToCommentResponse(comment)
}

func (service *CommentServiceImpl) Delete(ctx context.Context, commentId int) {
	err := service.Validate.Var(commentId, "required")
	helper.PanicIfError(err)

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		category, err := service.CommentRepository.FindById(ctx, tx, commentId)
		if err != nil {
			panic(exception.NewNotFoundError("Comment not found"))
		}

		err = service.CommentRepository.Delete(ctx, tx, category)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)
}
