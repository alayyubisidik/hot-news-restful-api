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

// import (
// 	"hot_news_2/helper"
// 	"hot_news_2/model/web"
// 	"net/http"
// 	"strings"
// )

// type AuthMiddleware struct {
// 	Handler http.Handler
// }

// func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
// 	return &AuthMiddleware{Handler: handler}
// }

// func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	path := r.URL.Path
// 	method := r.Method

// 	authRoutes := map[string][]string{
// 		"GET": {

// 		},
// 		"POST": {
// 			"/api/v1/categories",
// 			"/api/v1/articles",
// 			"/api/v1/comments",
// 			"/api/v1/likes",
// 		},
// 		"DELETE": {
// 			"/api/v1/users/signout",
// 			"/api/v1/categories",
// 			"/api/v1/articles",
// 			"/api/v1/comments",
// 			"/api/v1/likes",
// 		},
// 		"PUT": {
// 			"/api/v1/users",
// 			"/api/v1/categories",
// 			"/api/v1/articles",
// 			"/api/v1/comments",
// 		},
// 		// Tambahkan metode dan rute lainnya sesuai kebutuhan
// 	}

// 	requireAuth := false
// 	if routes, ok := authRoutes[method]; ok {
// 		for _, route := range routes {
// 			if strings.HasPrefix(path, route) {
// 				requireAuth = true
// 				break
// 			}
// 		}
// 	}

// 	if requireAuth {
// 		tokenCookie, err := r.Cookie("jwt")
// 		if err != nil {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusUnauthorized)

// 			webResponse := web.ErrorResponse{
// 				Errors: []web.DetailError{
// 					{
// 						Message: "Unauthorized",
// 					},
// 				},
// 			}

// 			helper.WriteToResponseBody(w, webResponse, 401)
// 			return
// 		}

// 		tokenString := tokenCookie.Value
// 		_, err = helper.VerifyToken(tokenString)
// 		if err != nil {
// 			w.Header().Set("Content-Type", "application/json")
// 			w.WriteHeader(http.StatusUnauthorized)

// 			webResponse := web.ErrorResponse{
// 				Errors: []web.DetailError{
// 					{
// 						Message: "Unauthorized",
// 					},
// 				},
// 			}

// 			helper.WriteToResponseBody(w, webResponse, 401)
// 			return
// 		}
// 	}

// 	middleware.Handler.ServeHTTP(w, r)
// }