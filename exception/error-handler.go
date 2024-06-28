package exception

import (
	"hot_news_2/helper"
	"hot_news_2/model/web"
	"net/http"

	"github.com/go-playground/validator"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err any) {
	if notFoundError(writer, request, err) {
		return
	}

	if validationError(writer, request, err) {
		return
	}

	if badRequestError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.WebResponse{
			Code: http.StatusNotFound,
			Status: "NOT FOUND",
			Data: exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: exception.Error(),
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func badRequestError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.WebResponse{
			Code: http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data: exception.Error,
		}

		helper.WriteToResponseBody(writer, webResponse)

		return true
	} else {
		return false
	}
}

func internalServerError(writer http.ResponseWriter, request *http.Request, err any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	var errorMessage string
    if e, ok := err.(error); ok {
        errorMessage = e.Error()
    } else {
        errorMessage = "Unknown error"
    }

    webResponse := web.WebResponse{
        Code:   http.StatusInternalServerError,
        Status: "INTERNAL SERVER ERROR",
        Data:   errorMessage,
    }

	helper.WriteToResponseBody(writer, webResponse)
}