package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCategorySuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	category1 := CreateCategory(db, "Sport")
	category2 := CreateCategory(db, "Tech")

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/categories", nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	categories := responseBody["data"].([]any)

	categoryResponse1 := categories[0].(map[string]any)
	categoryResponse2 := categories[1].(map[string]any)

	assert.Equal(t, category1.Name, categoryResponse1["name"])
	assert.Equal(t, category1.Slug, categoryResponse1["slug"])

	assert.Equal(t, category2.Name, categoryResponse2["name"])
	assert.Equal(t, category2.Slug, categoryResponse2["slug"])
}

func TestGetCategoryBySlugSuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	category := CreateCategory(db, "Sport")

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/categories/show/" + category.Slug, nil)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, category.Name, responseBody["data"].(map[string]any)["name"])
	assert.Equal(t, category.Slug, responseBody["data"].(map[string]any)["slug"])
}

func TestGetCategoryBySlugFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	category := CreateCategory(db, "Sport")

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/v1/categories/show/test" + category.Slug, nil)
	request.Header.Add("Content-Type", "application/json")

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}

func TestCreateCategorySuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name": "Sport"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 201, response.StatusCode)
	fmt.Println(response.Body)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "Sport", responseBody["data"].(map[string]any)["name"])
	assert.Equal(t, "sport", responseBody["data"].(map[string]any)["slug"])
}
 
func TestCreateCategoryFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/v1/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestUpdateCategorySuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	category := CreateCategory(db, "Sport")

	requestBody := strings.NewReader(`{"name": "Technology"}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/categories/" + category.Slug, requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]any
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, "Technology", responseBody["data"].(map[string]any)["name"])
	assert.Equal(t, "technology", responseBody["data"].(map[string]any)["slug"])
}

func TestUpdateCategoryFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	category := CreateCategory(db, "Sport")

	requestBody := strings.NewReader(`{"name": ""}`)
	request := httptest.NewRequest(http.MethodPut, "http://localhost:3000/api/v1/categories/" + category.Slug, requestBody)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestDeleteCategorySuccess(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	category := CreateCategory(db, "Sport")

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/categories/" + category.Slug, nil)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)
}

func TestDeleteCategoryFailed(t *testing.T) {
	db := SetupTestDB()
	TruncateTables(db, "articles", "categories", "users")
	router := SetupRouter(db)

	category := CreateCategory(db, "Sport")

	request := httptest.NewRequest(http.MethodDelete, "http://localhost:3000/api/v1/categories/test" + category.Slug, nil)
	request.Header.Add("Content-Type", "application/json")

	AddJWTToCookie(request)

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 404, response.StatusCode)
}
