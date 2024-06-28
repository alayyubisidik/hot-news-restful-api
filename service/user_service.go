package service

import (
	"context"
	"hot_news_2/model/web"
)

type UserService interface {
	SignUp(ctx context.Context, request web.UserSignUpRequest) web.AuthResponse
	SignIn(ctx context.Context, request web.UserSignInRequest) web.AuthResponse
	Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse
}