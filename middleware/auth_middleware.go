package middleware

import (
	"hot_news_2/helper"
	"hot_news_2/model/web"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path
	method := request.Method

	authRoutes := map[string][]string{
		"GET": {

		},
		"POST": {
			"/api/v1/categories",
			"/api/v1/articles",
			"/api/v1/comments",
		},
		"DELETE": {
			"/api/v1/users/signout",
			"/api/v1/categories",
			"/api/v1/articles",
			"/api/v1/comments",
		},
		"PUT": {
			"/api/v1/users",
			"/api/v1/categories",
			"/api/v1/articles",
			"/api/v1/comments",
		},
		// Tambahkan metode dan rute lainnya sesuai kebutuhan
	}

	requireAuth := false
	if routes, ok := authRoutes[method]; ok {
		for _, route := range routes {
			if strings.HasPrefix(path, route) {
				requireAuth = true
				break
			}
		}
	}

	if requireAuth {
		tokenCookie, err := request.Cookie("jwt")
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse)
			return
		}

		tokenString := tokenCookie.Value
		_, err = helper.VerifyToken(tokenString)
		if err != nil {
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse)
			return
		}
	}

	middleware.Handler.ServeHTTP(writer, request)
}