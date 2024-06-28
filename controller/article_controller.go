package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type ArticleController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByCategory(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindByUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindBySlug(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}