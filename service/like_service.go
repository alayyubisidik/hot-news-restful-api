package service

import (
	"context"
	"hot_news_2/model/web"
)

type LikeService interface {
	FindById(ctx context.Context, likeId int) web.LikeResponse
	Create(ctx context.Context, request web.LikeCreateRequest) web.LikeResponse
	Delete(ctx context.Context, likeId int)
} 