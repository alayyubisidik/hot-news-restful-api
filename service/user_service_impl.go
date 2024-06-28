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

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *gorm.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, db *gorm.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             db,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) SignUp(ctx context.Context, request web.UserSignUpRequest) web.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	hashedPassword, err := helper.HashPassword(request.Password)
	helper.PanicIfError(err)

	user := domain.User{
		Username: request.Username,
		FullName: request.FullName,
		Email:    request.Email,
		Password: hashedPassword,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		result, err := service.UserRepository.FindByUsername(ctx, tx, user.Username)
		if err == nil && result.ID != 0 {
			panic(exception.NewBadRequestError("Username is already exists"))
		}

		result, err = service.UserRepository.FindByEmail(ctx, tx, user.Email)
		if err == nil && result.ID != 0 {
			panic(exception.NewBadRequestError("Email is already exists"))
		}

		user, err = service.UserRepository.Create(ctx, tx, user)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	jwtToken, err := helper.CreateToken(user)
	helper.PanicIfError(err)

	return web.AuthResponse{
		Id: user.ID,
		Username: request.Username,
		FullName: request.FullName,
		Email:    request.Email,
		Token: jwtToken,
	}
}

func (service *UserServiceImpl) SignIn(ctx context.Context, request web.UserSignInRequest) web.AuthResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindByUsername(ctx, service.DB, request.Username)
	if err != nil {
		panic(exception.NewBadRequestError("Invalid credentials"))
	}
	
	err = helper.ComparePassword(user.Password, request.Password)
	if err != nil {
		panic(exception.NewBadRequestError("Invalid credentials"))
	}
	
	jwtToken, err := helper.CreateToken(user)
	helper.PanicIfError(err)

	return web.AuthResponse{
		Id: user.ID,
		Username: user.Username,
		FullName: user.FullName,
		Email:    user.Email,
		Token: jwtToken,
	}
}

func (service *UserServiceImpl) Update(ctx context.Context, request web.UserUpdateRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	user := domain.User{
		ID: request.Id,
		Username: request.Username,
		FullName: request.FullName,
		Email: request.Email,
	}

	err = service.DB.Transaction(func(tx *gorm.DB) error {
		_, err := service.UserRepository.FindById(ctx, tx, user.ID)
		if err != nil {
			panic(exception.NewNotFoundError("User not found"))
		}
		
		result, err := service.UserRepository.FindByUsername(ctx, tx, request.Username)
        if err == nil && result.ID != 0 && result.ID != user.ID {
            panic(exception.NewBadRequestError("Username already exists"))
        }

        result, err = service.UserRepository.FindByEmail(ctx, tx, request.Email)
        if err == nil && result.ID != 0 && result.ID != user.ID {
            panic(exception.NewBadRequestError("Email already exists"))
        }

		user, err = service.UserRepository.Update(ctx, tx, user)
		helper.PanicIfError(err)

		return nil
	})

	helper.PanicIfError(err)

	return helper.ToUserResponse(user)
}
