package app

import (
	"hot_news_2/controller"
	"hot_news_2/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController, categoryController controller.CategoryController, articleController controller.ArticleController, commentController controller.CommentController) *httprouter.Router {
    router := httprouter.New()

    router.POST("/api/v1/users/signup", userController.SignUp)
    router.POST("/api/v1/users/signin", userController.SignIn)
    router.DELETE("/api/v1/users/signout", userController.SignOut)
    router.GET("/api/v1/users/currentuser", userController.CurrentUser)
    router.PUT("/api/v1/users/:userId", userController.Update)

    router.GET("/api/v1/categories", categoryController.FindAll)
    router.GET("/api/v1/categories/show/:categorySlug", categoryController.FindBySlug)
    router.POST("/api/v1/categories", categoryController.Create)
    router.PUT("/api/v1/categories/:categorySlug", categoryController.Update)
    router.DELETE("/api/v1/categories/:categorySlug", categoryController.Delete)

    router.GET("/api/v1/articles", articleController.FindAll)
	router.GET("/api/v1/articles/categories/:categorySlug", articleController.FindByCategory)
    router.GET("/api/v1/articles/users/:username", articleController.FindByUser)
    router.GET("/api/v1/articles/show/:articleSlug", articleController.FindBySlug)
    router.POST("/api/v1/articles", articleController.Create)
    router.PUT("/api/v1/articles/:articleSlug", articleController.Update)
    router.DELETE("/api/v1/articles/:articleSlug", articleController.Delete)

    router.GET("/api/v1/comments/users/:username", commentController.FindByUser)
    router.GET("/api/v1/comments/show/:commentId", commentController.FindById)
    router.POST("/api/v1/comments", commentController.Create)
    router.PUT("/api/v1/comments/:commentId", commentController.Update)
    router.DELETE("/api/v1/comments/:commentId", commentController.Delete)

    router.PanicHandler = exception.ErrorHandler

    return router
}



