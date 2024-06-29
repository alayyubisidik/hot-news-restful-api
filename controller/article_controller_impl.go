package controller

import (
	"hot_news_2/exception"
	"hot_news_2/helper"
	"hot_news_2/model/web"
	"hot_news_2/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)
 
type ArticleControllerImpl struct {
	ArticleService service.ArticleService
}

func NewArticleController(articleService service.ArticleService) ArticleController {
	return &ArticleControllerImpl{
		ArticleService: articleService,
	}
}

func (controller *ArticleControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	articlesResponse := controller.ArticleService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: articlesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArticleControllerImpl) FindByCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categorySlug := params.ByName("categorySlug")

	articleResponses := controller.ArticleService.FindByCategory(request.Context(), categorySlug)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: articleResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArticleControllerImpl) FindByUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	username := params.ByName("username")

	articleResponses := controller.ArticleService.FindByUser(request.Context(), username)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: articleResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArticleControllerImpl) FindBySlug(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	articleSlug := params.ByName("articleSlug")

	articleResponse := controller.ArticleService.FindBySlug(request.Context(), articleSlug)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: articleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArticleControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	articleCreateRequest := web.ArticleCreateRequest{}
	err := helper.ReadFromRequestBody(request, &articleCreateRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	articleResponse := controller.ArticleService.Create(request.Context(), articleCreateRequest)
	webResponse := web.WebResponse{
		Code: 201,
		Status: "OK",
		Data: articleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArticleControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	articleUpdateRequest := web.ArticleUpdateRequest{}
	err := helper.ReadFromRequestBody(request, &articleUpdateRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	articleSlug := params.ByName("articleSlug")

	articleResponse := controller.ArticleService.Update(request.Context(), articleUpdateRequest, articleSlug)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: articleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *ArticleControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	articleSlug := params.ByName("articleSlug")

	controller.ArticleService.Delete(request.Context(), articleSlug)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}