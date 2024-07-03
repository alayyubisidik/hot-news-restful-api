package app

import (
	"hot_news_2/controller"
	"hot_news_2/exception"
	"hot_news_2/middleware"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(userController controller.UserController, categoryController controller.CategoryController, articleController controller.ArticleController, commentController controller.CommentController, likeController controller.LikeController) *httprouter.Router {
    router := httprouter.New()

    router.POST("/api/v1/users/signup", userController.SignUp)
    router.POST("/api/v1/users/signin", userController.SignIn)
    router.DELETE("/api/v1/users/signout", middleware.ChainMiddleware(userController.SignOut, middleware.AuthMiddleware))
    router.GET("/api/v1/users/currentuser", middleware.ChainMiddleware(userController.CurrentUser, middleware.AuthMiddleware))
    router.PUT("/api/v1/users/:userId", middleware.ChainMiddleware(userController.Update, middleware.AuthMiddleware))

    router.GET("/api/v1/categories", categoryController.FindAll)
    router.GET("/api/v1/categories/show/:categorySlug", categoryController.FindBySlug)
    router.POST("/api/v1/categories", middleware.ChainMiddleware(categoryController.Create, middleware.AuthMiddleware))
    router.PUT("/api/v1/categories/:categorySlug", middleware.ChainMiddleware(categoryController.Update, middleware.AuthMiddleware))
    router.DELETE("/api/v1/categories/:categorySlug", middleware.ChainMiddleware( categoryController.Delete, middleware.AuthMiddleware))

    router.GET("/api/v1/articles", articleController.FindAll)
	router.GET("/api/v1/articles/categories/:categorySlug", articleController.FindByCategory)
    router.GET("/api/v1/articles/users/:username", articleController.FindByUser)
    router.GET("/api/v1/articles/show/:articleSlug", articleController.FindBySlug)
    router.POST("/api/v1/articles", middleware.ChainMiddleware(articleController.Create, middleware.AuthMiddleware))
    router.PUT("/api/v1/articles/:articleSlug", middleware.ChainMiddleware(articleController.Update, middleware.AuthMiddleware))
    router.DELETE("/api/v1/articles/:articleSlug", middleware.ChainMiddleware(articleController.Delete, middleware.AuthMiddleware))
 
    router.GET("/api/v1/comments/users/:username", commentController.FindByUser)
    router.GET("/api/v1/comments/show/:commentId", commentController.FindById)
    router.POST("/api/v1/comments", middleware.ChainMiddleware(commentController.Create , middleware.AuthMiddleware))
    router.PUT("/api/v1/comments/:commentId", middleware.ChainMiddleware(commentController.Update , middleware.AuthMiddleware))
    router.DELETE("/api/v1/comments/:commentId", middleware.ChainMiddleware(commentController.Delete , middleware.AuthMiddleware))

    router.GET("/api/v1/likes/show/:likeId", likeController.FindById)
    router.POST("/api/v1/likes", middleware.ChainMiddleware(likeController.Create , middleware.AuthMiddleware))
    router.DELETE("/api/v1/likes/:likeId", middleware.ChainMiddleware(likeController.Delete , middleware.AuthMiddleware))

    router.PanicHandler = exception.ErrorHandler

    return router 
}



 