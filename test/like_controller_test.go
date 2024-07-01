package test

import (
	"encoding/json"
	"strconv"
	"strings"

	// "fmt"
	"io"
	"net/http"

	// "strings"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestGetLikeByIdSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments", "likes")
	router := SetupRouter(db)

	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	article := CreateArticle(db, "Title", user.ID, category.ID)

	like := CreateLike(db, user.ID, article.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/likes/show/" + strconv.Itoa(like.ID), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, float64(like.ID), responseBody["data"].(map[string]any)["id"])
}

func TestGetLikeByIdFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments", "likes")
	router := SetupRouter(db)

	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	article := CreateArticle(db, "Title", user.ID, category.ID)

	like := CreateLike(db, user.ID, article.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/likes/show/" + strconv.Itoa(like.ID  + 1), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestCreateLikeSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments", "likes")
	router := SetupRouter(db)
	
	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	CreateArticle(db, "Title", user.ID, category.ID)

	requestBody := strings.NewReader(`{"user_id": 1, "article_id": 1}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/likes", requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, float64(1), responseBody["data"].(map[string]any)["id"])
}

func TestCreateLikeFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments", "likes")
	router := SetupRouter(db)
	
	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	CreateArticle(db, "Title", user.ID, category.ID)

	requestBody := strings.NewReader(`{"user_id": 1, "article_id": 1}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/likes", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}

func TestDeleteLikeSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments", "likes")
	router := SetupRouter(db)

	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	article := CreateArticle(db, "Title", user.ID, category.ID)

	like := CreateLike(db, user.ID, article.ID)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/likes/" + strconv.Itoa(like.ID), nil)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteLikeFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments", "likes")
	router := SetupRouter(db)

	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	article := CreateArticle(db, "Title", user.ID, category.ID)

	like := CreateLike(db, user.ID, article.ID)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/likes/" + strconv.Itoa(like.ID + 1), nil)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}