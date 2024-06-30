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

type LikeControllerImpl struct {
	LikeService service.LikeService
}

func NewLikeController(likeService service.LikeService) LikeController {
	return &LikeControllerImpl{
		LikeService: likeService,
	}
}

func (controller *LikeControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	likeId := params.ByName("likeId")
	id, err := strconv.Atoi(likeId)
	helper.PanicIfError(err)

	likeResponse := controller.LikeService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Data: likeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}

func (controller *LikeControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	likeCreateRequest := web.LikeCreateRequest{}
	err := helper.ReadFromRequestBody(request, &likeCreateRequest)
	if err != nil {
		panic(exception.NewBadRequestError(err.Error()))
	}

	likeResponse := controller.LikeService.Create(request.Context(), likeCreateRequest)
	webResponse := web.WebResponse{
		Data: likeResponse,
	}

	helper.WriteToResponseBody(writer, webResponse, 201)
}

func (controller *LikeControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	likeId := params.ByName("likeId")
	id, err := strconv.Atoi(likeId)
	helper.PanicIfError(err)


	controller.LikeService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Data: "Like successfully deleted",
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}
