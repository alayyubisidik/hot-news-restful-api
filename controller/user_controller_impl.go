package controller

import (
	"hot_news_2/helper"
	"hot_news_2/model/web"
	"hot_news_2/service"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) SignUp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userSignUpRequest := web.UserSignUpRequest{}
	helper.ReadFromRequestBody(request, &userSignUpRequest)

	authResponse := controller.UserService.SignUp(request.Context(), userSignUpRequest)

    http.SetCookie(writer, &http.Cookie{
        Name:     "jwt",
        Value:    authResponse.Token,
        Path:     "/",
        HttpOnly: true,   
        SameSite: http.SameSiteStrictMode, 
    })

    webResponse := web.WebResponse{
		Data: struct {
			Id        int       `json:"id"`
			Username  string    `json:"username"`
			FullName  string    `json:"full_name"`
			Email     string    `json:"email"`
			CreatedAt time.Time `json:"created_at"`
		}{
			Id:        authResponse.Id,
			Username:  authResponse.Username,
			FullName:  authResponse.FullName,
			Email:     authResponse.Email,
			CreatedAt: authResponse.CreatedAt,
		},
    }

    helper.WriteToResponseBody(writer, webResponse, 201)
}

func (controller *UserControllerImpl) SignIn(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userSignInRequest := web.UserSignInRequest{}
	helper.ReadFromRequestBody(request, &userSignInRequest)

	authResponse := controller.UserService.SignIn(request.Context(), userSignInRequest)

    http.SetCookie(writer, &http.Cookie{
        Name:     "jwt",
        Value:    authResponse.Token,
        Path:     "/",
        HttpOnly: true,   
        SameSite: http.SameSiteStrictMode, 
    })
 
    webResponse := web.WebResponse{
		Data: struct {
			Id        int       `json:"id"`
			Username  string    `json:"username"`
			FullName  string    `json:"full_name"`
			Email     string    `json:"email"`
			CreatedAt time.Time `json:"created_at"`
		}{
			Id:        authResponse.Id,
			Username:  authResponse.Username,
			FullName:  authResponse.FullName,
			Email:     authResponse.Email,
			CreatedAt: authResponse.CreatedAt,
		},
    }

    helper.WriteToResponseBody(writer, webResponse, 200)
}

func (controller *UserControllerImpl) SignOut(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	http.SetCookie(writer, &http.Cookie{
		Name:     "jwt",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
	})

	webResponse := web.WebResponse{
		Data: "Signout successfully",
	}

	helper.WriteToResponseBody(writer, webResponse, 200)
}

func (controller *UserControllerImpl) CurrentUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
    tokenCookie, err := request.Cookie("jwt")
    if err != nil {
        writer.Header().Set("Content-Type", "application/json")
        webResponse := web.ErrorResponse{
			Errors: []web.DetailError{
				{
					Message: "Unauthorized",
				},
			},
        }
        helper.WriteToResponseBody(writer, webResponse, 401)
        return
    }

    tokenString := tokenCookie.Value

    claims, err := helper.VerifyToken(tokenString)
    if err != nil {
        writer.Header().Set("Content-Type", "application/json")
        webResponse := web.ErrorResponse{
			Errors: []web.DetailError{
				{
					Message: "Unauthorized",
				},
			},
        }
        helper.WriteToResponseBody(writer, webResponse, 401)
        return
    }

    userResponse := web.UserResponse{
        Id:       claims.ID,
        Username: claims.Username,
        FullName: claims.FullName,
        Email:    claims.Email,
    }

    webResponse := web.WebResponse{
        Data:   userResponse,
    }

    helper.WriteToResponseBody(writer, webResponse, 200)
}


func (controller *UserControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userUpdateRequest := web.UserUpdateRequest{}
	helper.ReadFromRequestBody(request, &userUpdateRequest)

	userId := params.ByName("userId")
	id, err := strconv.Atoi(userId)
	helper.PanicIfError(err)

	userUpdateRequest.Id = id

	userResponse := controller.UserService.Update(request.Context(), userUpdateRequest)
    webResponse := web.WebResponse{
        Data:   userResponse,
    }

	helper.WriteToResponseBody(writer, webResponse, 200)
}