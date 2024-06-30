package controller

import (
	"hot_news_2/exception"
	"hot_news_2/helper"
	"hot_news_2/model/web"
	"hot_news_2/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CommentControllerImpl struct {
	CommentService service.CommentService
}

func NewCommentController(commentService service.CommentService) CommentController {
	return &CommentControllerImpl{
		CommentService: commentService,
	}
}

func (controller *CommentControllerImpl) FindByUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	username := params.ByName("username")

	commentResponses := controller.CommentService.FindByUser(request.Context(), username)
	webResponse := web.WebResponse{
		Data: commentResponses,
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
} 

func (controller *CommentControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	commentId := params.ByName("commentId")
	id, err := strconv.Atoi(commentId)
	helper.PanicIfError(err)

	commentResponse := controller.CommentService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Data: commentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
} 

func (controller *CommentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	commentCreateRequest := web.CommentCreateRequest{}
	err := helper.ReadFromRequestBody(request, &commentCreateRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	commentResponse := controller.CommentService.Create(request.Context(), commentCreateRequest)
	webResponse := web.WebResponse{
		Data: commentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 201)
} 

func (controller *CommentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	commentUpdateRequest := web.CommentUpdateRequest{}
	err := helper.ReadFromRequestBody(request, &commentUpdateRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	commentId := params.ByName("commentId")
	id, err := strconv.Atoi(commentId)
	helper.PanicIfError(err)

	commentUpdateRequest.Id = id

	commentResponse := controller.CommentService.Update(request.Context(), commentUpdateRequest)
	webResponse := web.WebResponse{
		Data: commentResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
} 

func (controller *CommentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	commentId := params.ByName("commentId")
	id, err := strconv.Atoi(commentId)
	helper.PanicIfError(err)


	controller.CommentService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Data: "Comment successfully deleted",
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
} 