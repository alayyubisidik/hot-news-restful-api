package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CategoryController interface {
	FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	FindBySlug(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}