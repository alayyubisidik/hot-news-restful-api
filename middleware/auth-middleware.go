package middleware

import (
	"hot_news_2/helper"
	"hot_news_2/model/web"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func AuthMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		tokenCookie, err := r.Cookie("jwt")
		if err != nil {
			w.Header().Set("Content-Type", "application/json")

			webResponse := web.ErrorResponse{
				Errors: []web.DetailError{
					{
						Message: "Unauthorized",
					},
				},
			}

			helper.WriteToResponseBody(w, webResponse, 401)
			return
		}

		tokenString := tokenCookie.Value
		_, err = helper.VerifyToken(tokenString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")

			webResponse := web.ErrorResponse{
				Errors: []web.DetailError{
					{
						Message: "Unauthorized",
					},
				},
			}

			helper.WriteToResponseBody(w, webResponse, 401)
			return
		}

		next(w, r, ps)
	}
}

