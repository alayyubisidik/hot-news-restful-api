package controller

import (
	"hot_news_2/helper"
	"hot_news_2/model/web"
	"hot_news_2/service"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CategoryController {
	return &CategoryControllerImpl{
		CategoryService: categoryService,
	}
}

func (controller *CategoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoriesResponse := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Data: categoriesResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}

func (controller *CategoryControllerImpl) FindBySlug(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categorySlug := params.ByName("categorySlug")

	categoryResponse := controller.CategoryService.FindBySlug(request.Context(), categorySlug)
	webResponse := web.WebResponse{
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}

func (controller *CategoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)

	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 201)
}

func (controller *CategoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categorySlug := params.ByName("categorySlug")

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest, categorySlug)
	webResponse := web.WebResponse{
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}

func (controller *CategoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	categorySlug := params.ByName("categorySlug")

	controller.CategoryService.Delete(request.Context(), categorySlug)
	webResponse := web.WebResponse{
		Data: "Category successfully deleted",
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}