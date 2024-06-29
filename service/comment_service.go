package service

import (
	"context"
	"hot_news_2/model/web"
)

type CommentService interface {
	FindByUser(ctx context.Context, username string) []web.CommentResponse
	FindById(ctx context.Context, commentId int) web.CommentResponse
	Create(ctx context.Context, request web.CommentCreateRequest) web.CommentResponse
	Update(ctx context.Context, request web.CommentUpdateRequest) web.CommentResponse
	Delete(ctx context.Context, commentInt int)
} 