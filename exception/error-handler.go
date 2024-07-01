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

	if unauthorizedError(writer, request, err) {
		return
	}

	internalServerError(writer, request, err)
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.ErrorResponse{
			Errors: []web.DetailError{
				{
					Message: exception.Error,
				},
			},
		}

		helper.WriteToResponseBody(writer, webResponse, 404)

		return true
	} else {
		return false
	}
}

func validationError(writer http.ResponseWriter, request *http.Request, err any) bool {
	validationErr, ok := err.(validator.ValidationErrors)
	if !ok {
		return false
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusBadRequest)

	var errors []web.DetailError

	for _, fieldError := range validationErr {
		errorDetail := web.DetailError{
			Field:   fieldError.Field(),
			Message: fieldError.Tag(),
		}
		errors = append(errors, errorDetail)
	}
 
	webResponse := web.ErrorResponse{
		Errors: errors,
	}

	helper.WriteToResponseBody(writer, webResponse, http.StatusBadRequest)

	return true
}


func badRequestError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		
		webResponse := web.ErrorResponse{
			Errors: []web.DetailError{
				{
					Message: exception.Error,
				},
			},
		}

		helper.WriteToResponseBody(writer, webResponse, 400)

		return true
	} else {
		return false
	}
}

func unauthorizedError(writer http.ResponseWriter, request *http.Request, err any) bool {
	exception, ok := err.(UnauthorizedError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		webResponse := web.ErrorResponse{
			Errors: []web.DetailError{
				{
					Message: exception.Error,
				},
			},
		}

		helper.WriteToResponseBody(writer, webResponse, 404)

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

	webResponse := web.ErrorResponse{
		Errors: []web.DetailError{
			{
				Message: errorMessage,
			},
		},
	}

	helper.WriteToResponseBody(writer, webResponse, 500)
}