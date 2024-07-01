package test

import (
	"encoding/json"
	"strings"

	// "fmt"
	"io"
	"net/http"

	// "strings"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestGetAllArticleSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	category := CreateCategory(db, "Sport")
	user := CreateUser(db, "test", "test@gmail.com")

	article1 := CreateArticle(db, "Title1", user.ID, category.ID)
	article2 := CreateArticle(db, "Title2", user.ID, category.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/articles", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	articles := responseBody["data"].([]any)

	articleResponse1 := articles[0].(map[string]any)
	articleResponse2 := articles[1].(map[string]any)

	assert.Equal(t, article1.Title, articleResponse1["title"])
	assert.Equal(t, article1.Slug, articleResponse1["slug"])

	assert.Equal(t, article2.Title, articleResponse2["title"])
	assert.Equal(t, article2.Slug, articleResponse2["slug"])
}

func TestGetArticleByCategorySuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	category1 := CreateCategory(db, "Sport")
	category2 := CreateCategory(db, "Technology")
	user := CreateUser(db, "test", "test@gmail.com")

	article1 := CreateArticle(db, "Title1", user.ID, category1.ID)
	article2 := CreateArticle(db, "Title2", user.ID, category1.ID)
	CreateArticle(db, "Title3", user.ID, category2.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/articles/categories/"+category1.Slug, nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	articles := responseBody["data"].([]any)

	assert.Equal(t, 2, len(articles))

	articleResponse1 := articles[0].(map[string]any)
	articleResponse2 := articles[1].(map[string]any)

	categoryResponse1 := articleResponse1["category"].(map[string]any)
	categoryResponse2 := articleResponse2["category"].(map[string]any)

	assert.Equal(t, article1.Title, articleResponse1["title"])
	assert.Equal(t, article1.Slug, articleResponse1["slug"])
	assert.Equal(t, float64(category1.ID), categoryResponse1["id"])

	assert.Equal(t, article2.Title, articleResponse2["title"])
	assert.Equal(t, article2.Slug, articleResponse2["slug"])
	assert.Equal(t, float64(category1.ID), categoryResponse2["id"])

}

func TestGetArticleByCategoryFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	category1 := CreateCategory(db, "Sport")
	category2 := CreateCategory(db, "Technology")
	user := CreateUser(db, "test", "test@gmail.com")

	CreateArticle(db, "Title1", user.ID, category1.ID)
	CreateArticle(db, "Title2", user.ID, category1.ID)
	CreateArticle(db, "Title3", user.ID, category2.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/articles/categories/notfound"+category1.Slug, nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestGetArticleByUserSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	user1 := CreateUser(db, "user1", "user1@gmail.com")
	user2 := CreateUser(db, "user2", "user2@gmail.com")
	category := CreateCategory(db, "Sport")

	article1 := CreateArticle(db, "Title1", user1.ID, category.ID)
	article2 := CreateArticle(db, "Title2", user1.ID, category.ID)
	CreateArticle(db, "Title3", user2.ID, category.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/articles/users/" + user1.Username, nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	articles := responseBody["data"].([]any)

	articleResponse1 := articles[0].(map[string]any)
	articleResponse2 := articles[1].(map[string]any)

	userResponse1 := articleResponse1["user"].(map[string]any)
	userResponse2 := articleResponse2["user"].(map[string]any)

	assert.Equal(t, article1.Title, articleResponse1["title"])
	assert.Equal(t, article1.Slug, articleResponse1["slug"])
	assert.Equal(t, float64(user1.ID), userResponse1["id"])

	assert.Equal(t, article2.Title, articleResponse2["title"])
	assert.Equal(t, article2.Slug, articleResponse2["slug"])
	assert.Equal(t, float64(user1.ID), userResponse2["id"])
}

func TestGetArticleByUserFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user1 := CreateUser(db, "user1", "user1@gmail.com")
	user2 := CreateUser(db, "user2", "user2@gmail.com")
	category := CreateCategory(db, "Sport")

	CreateArticle(db, "Title1", user1.ID, category.ID)
	CreateArticle(db, "Title2", user1.ID, category.ID)
	CreateArticle(db, "Title3", user2.ID, category.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/articles/users/notfound" + user1.Username, nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestGetArticleBySlugSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user1", "user1@gmail.com")
	category := CreateCategory(db, "Sport")

	article := CreateArticle(db, "Title1", user.ID, category.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/articles/show/" + article.Slug, nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, article.Title, responseBody["data"].(map[string]any)["title"])
	assert.Equal(t, article.Slug, responseBody["data"].(map[string]any)["slug"])
}

func TestGetArticleBySlugFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user1", "user1@gmail.com")
	category := CreateCategory(db, "Sport")

	article := CreateArticle(db, "Title1", user.ID, category.ID)

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/articles/show/notfound" + article.Slug, nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestCreateArticleSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	CreateUser(db, "user1", "user1@gmail.com")
	CreateCategory(db, "Sport")

	requestBody := strings.NewReader(`{
		"title": "Article 1",
		"content": "loremblablablcablnaca",
		"user_id": 1,
		"category_id": 1
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/articles", requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "article-1", responseBody["data"].(map[string]any)["slug"])
}

func TestCreateArticleFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	CreateUser(db, "user1", "user1@gmail.com")
	CreateCategory(db, "Sport")

	requestBody := strings.NewReader(`{
		"titl": "Article 1",
		"content": "loremblablablcablnaca",
		"user_id": 1,
		"category_id": 1
	}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/articles", requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}


func TestUpdateArticleSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user1", "user1@gmail.com")
	category := CreateCategory(db, "Sport")

	article := CreateArticle(db, "Title1", user.ID, category.ID)

	requestBody := strings.NewReader(`{
		"title": "Article 2",
		"content": "loremblablablcablnaca",
		"category_id": 1
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/articles/" + article.Slug, requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "article-2", responseBody["data"].(map[string]any)["slug"])
}

func TestUpdateArticleFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user1", "user1@gmail.com")
	category := CreateCategory(db, "Sport")

	article := CreateArticle(db, "Title1", user.ID, category.ID)

	requestBody := strings.NewReader(`{
		"title": "Article 2",
		"content": "loremblablablcablnaca",
		"category_id": 1
	}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/articles/notfound" + article.Slug, requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestDeleteArticleSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user1", "user1@gmail.com")
	category := CreateCategory(db, "Sport")

	article := CreateArticle(db, "Title1", user.ID, category.ID)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/articles/" + article.Slug, nil)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteArticleFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users", "comments")
	router := SetupRouter(db)

	user := CreateUser(db, "user1", "user1@gmail.com")
	category := CreateCategory(db, "Sport")

	article := CreateArticle(db, "Title1", user.ID, category.ID)

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/articles/notfound" + article.Slug, nil)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}