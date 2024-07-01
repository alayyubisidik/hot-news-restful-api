package test

import (
	"context"
	"hot_news_2/app"
	"hot_news_2/controller"
	"hot_news_2/helper"
	"hot_news_2/middleware"
	"hot_news_2/model/domain"
	"hot_news_2/repository"
	"hot_news_2/service"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gosimple/slug"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/hot_news_2_test?charset=utf8mb4&parseTime=True&loc=Local"))
	helper.PanicIfError(err)

	return db
}

func SetupRouter(db *gorm.DB) http.Handler {
	userRepository := repository.NewUserRepository()
	validate := validator.New()
	userService := service.NewUserService(userRepository, db, validate)
	userController := controller.NewUserController(userService)
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	articleRepository := repository.NewArticleRepository()
	articleService := service.NewArticleService(articleRepository, db, validate)
	articleController := controller.NewArticleController(articleService)
	commentRepository := repository.NewCommentRepository()
	commentService := service.NewCommentService(commentRepository, db, validate)
	commentController := controller.NewCommentController(commentService)
	likeRepository := repository.NewLikeRepository()
	likeService := service.NewLikeService(likeRepository, db, validate)
	likeController := controller.NewLikeController(likeService)
	router := app.NewRouter(userController, categoryController, articleController, commentController, likeController)
	return middleware.NewAuthMiddleware(router)
}

func AddJWTToCookie(request *http.Request) {
	user := domain.User{
		ID:       1,
		Username: "test",
		FullName: "Test",
		Email:    "test@gmail.com",
		Password: "password",
	}

	jwtToken, err := helper.CreateToken(user)
	if err != nil {
		helper.PanicIfError(err)
	}

	cookie := &http.Cookie{
		Name:     "jwt",
		Value:    jwtToken,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	}
	request.AddCookie(cookie)
}

func TruncateTable(db *gorm.DB, tableName string) error {
	sql := "TRUNCATE TABLE " + tableName
	return db.Exec(sql).Error
}

func TruncateTables(db *gorm.DB, tables ...string) {
    db.Exec("SET FOREIGN_KEY_CHECKS = 0;")
    for _, table := range tables {
        db.Exec("TRUNCATE TABLE " + table)
    }
    db.Exec("SET FOREIGN_KEY_CHECKS = 1;")
}

func CreateCategory(db *gorm.DB, name string) domain.Category {
	slug := slug.Make(name)

	category := domain.Category{
		Name: name,
		Slug: slug,
	}
	db.Transaction(func(tx *gorm.DB) error {
		var err error
		categoryRepository := repository.NewCategoryRepository()
		category, err = categoryRepository.Create(context.Background(), tx, category)
		helper.PanicIfError(err)

		return nil
	})
	
	return category
}

func CreateUser(db *gorm.DB, username string, email string) domain.User {
	hashedPassword, err := helper.HashPassword("test")
	helper.PanicIfError(err)

	user := domain.User{
		Username: username,
		FullName: "Test",
		Email: email,
		Password: hashedPassword,
	}

	db.Transaction(func(tx *gorm.DB) error {
		var err error
		userRepository := repository.NewUserRepository()
		user, err = userRepository.Create(context.Background(), tx, user)
		helper.PanicIfError(err)

		return nil
	})
	
	return user
}

func CreateArticle(db *gorm.DB, title string, userId int, categoryId int) domain.Article {
	slug := slug.Make(title)

	article := domain.Article{
		Title: title,
		Slug: slug,
		Content: "Lorem ipsum dolor sit amet consectetur adipisicing elit. Eius, excepturi cum quas doloremque ipsam fuga voluptates nulla animi officia tempora neque obcaecati maxime alias qui eos nemo ea quidem voluptas odit sint aspernatur voluptatem aut inventore suscipit. Libero qui commodi, dolores quaerat expedita neque repellendus! Sint odio ab sed cupiditate!",
		UserId: userId,
		CategoryId: categoryId,
	}
	db.Transaction(func(tx *gorm.DB) error {
		var err error
		articleRepository := repository.NewArticleRepository()
		article, err = articleRepository.Create(context.Background(), tx, article)
		helper.PanicIfError(err)

		return nil
	})
	
	return article
}

func CreateComment(db *gorm.DB, userId int, articleId int) domain.Comment {
	comment := domain.Comment{
		UserId: userId,
		ArticleId: articleId,
		Content: "loremwigowinvwrvrw",
	}
	db.Transaction(func(tx *gorm.DB) error {
		var err error
		commentRepository := repository.NewCommentRepository()
		comment, err = commentRepository.Create(context.Background(), tx, comment)
		helper.PanicIfError(err)

		return nil
	})
	
	return comment
}

func CreateLike(db *gorm.DB, userId int, articleId int) domain.Like {
	like := domain.Like{
		UserId: userId,
		ArticleId: articleId,
	}
	db.Transaction(func(tx *gorm.DB) error {
		var err error
		likeRepository := repository.NewLikeRepository()
		like, err = likeRepository.Create(context.Background(), tx, like)
		helper.PanicIfError(err)

		return nil
	})
	
	return like
}
