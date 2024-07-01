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

func TestGetCommentByUserSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user1 := CreateUser(db, "user1", "user1@gmail.com")
	user2 := CreateUser(db, "user2", "user2@gmail.com")
	category := CreateCategory(db, "Sport")

	article := CreateArticle(db, "Title1", user1.ID, category.ID)

	CreateComment(db, user1.ID, article.ID)
	CreateComment(db, user1.ID, article.ID)
	CreateComment(db, user2.ID, article.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/comments/users/" + user1.Username, nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	comments := responseBody["data"].([]any)

	commentResponse1 := comments[0].(map[string]any)
	commentResponse2 := comments[1].(map[string]any)

	userResponse1 := commentResponse1["user"].(map[string]any)
	userResponse2 := commentResponse2["user"].(map[string]any)

	assert.Equal(t, float64(user1.ID), userResponse1["id"])
	assert.Equal(t, float64(user1.ID), userResponse2["id"])	
}

func TestGetCommentByUserFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user1 := CreateUser(db, "user1", "user1@gmail.com")
	user2 := CreateUser(db, "user2", "user2@gmail.com")
	category := CreateCategory(db, "Sport")

	article := CreateArticle(db, "Title1", user1.ID, category.ID)

	CreateComment(db, user1.ID, article.ID)
	CreateComment(db, user1.ID, article.ID)
	CreateComment(db, user2.ID, article.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/comments/users/notfound" + user1.Username, nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestGetCommentByIdSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	article := CreateArticle(db, "Title", user.ID, category.ID)

	comment := CreateComment(db, user.ID, article.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/comments/show/" + strconv.Itoa(comment.ID), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, float64(comment.ID), responseBody["data"].(map[string]any)["id"])
}

func TestGetCommentByIdFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	article := CreateArticle(db, "Title", user.ID, category.ID)

	comment := CreateComment(db, user.ID, article.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/comments/show/" + strconv.Itoa(comment.ID + 1), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestCreateCommentSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)
	
	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	CreateArticle(db, "Title", user.ID, category.ID)

	requestBody := strings.NewReader(`{"user_id": 1, "article_id": 1, "content": "tetifauivbwvwevwvwvr"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/comments", requestBody)
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

func TestCreateCommentFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)
	
	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	CreateArticle(db, "Title", user.ID, category.ID)

	requestBody := strings.NewReader(`{"user_id": 1, "article_id": 1, "content": "tetifauivbwvwevwvwvr"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/comments", requestBody)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}

func TestUpdateCommentSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	article := CreateArticle(db, "Title", user.ID, category.ID)

	comment := CreateComment(db, user.ID, article.ID)

	requestBody := strings.NewReader(`{"content": "contentupdated"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/comments/" + strconv.Itoa(comment.ID), requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "contentupdated", responseBody["data"].(map[string]any)["content"])
}

func TestUpdateCommentFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	article := CreateArticle(db, "Title", user.ID, category.ID)

	comment := CreateComment(db, user.ID, article.ID)

	requestBody := strings.NewReader(`{"content": "contentupdated"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/comments/" + strconv.Itoa(comment.ID + 1), requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestDeleteCommentSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	article := CreateArticle(db, "Title", user.ID, category.ID)

	comment := CreateComment(db, user.ID, article.ID)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/comments/" + strconv.Itoa(comment.ID), nil)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteCommentFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user", "user@gmail.com")
	category := CreateCategory(db, "Sport")
	article := CreateArticle(db, "Title", user.ID, category.ID)

	comment := CreateComment(db, user.ID, article.ID)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/comments/" + strconv.Itoa(comment.ID), nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 401, response.StatusCode)
}