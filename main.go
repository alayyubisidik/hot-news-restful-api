package main

import (
	"hot_news_2/helper"
	"hot_news_2/middleware"
	"net/http"

)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr: "localhost:3000",
		Handler: authMiddleware,
	}
}

func main() {
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}