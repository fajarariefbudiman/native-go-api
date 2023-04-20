package test

import (
	"belajar-golang-api/app"
	"belajar-golang-api/controller"
	"belajar-golang-api/helper"
	"belajar-golang-api/middleware"
	"belajar-golang-api/model/domain"
	"belajar-golang-api/repository"
	"belajar-golang-api/service"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/go-playground/validator/v10"
)

func setupDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/category_repository_test")
	helper.PanicIfError(err)

	return db

}

func trancateCategory(db *sql.DB) {
	db.Exec("TRUNCATE customer")
}

func SetUpRouter(db *sql.DB) http.Handler {

	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)

	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	return middleware.NewAuthMiddleware(router)
}

func TestCreateCategorySuccess(t *testing.T) {
	db := setupDB()
	trancateCategory(db)
	router := SetUpRouter(db)

	requestbody := strings.NewReader(`{"name":"laptop"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestbody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

}

func TestCreateCategoryFailed(t *testing.T) {
	db := setupDB()
	trancateCategory(db)
	router := SetUpRouter(db)

	requestbody := strings.NewReader(`{"name":""}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestbody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 400, response.StatusCode)
}

func TestUpdateCategorySucces(t *testing.T) {
	db := setupDB()
	trancateCategory(db)

	tx, _ := db.Begin()
	categoryRepository := repository.NewCategoryRepository()
	category := categoryRepository.Save(context.Background(), tx, domain.Category{
		Name: "laptop",
	})

	tx.Commit()

	router := SetUpRouter(db)

	requestbody := strings.NewReader(`{"name":"laptop"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories/"+strconv.Itoa(category.Id), requestbody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responsebody map[string]interface{}
	json.Unmarshal(body, &responsebody)

}

func TestUpdateCategoryFailed(t *testing.T) {

}

func TestGetCategorySucces(t *testing.T) {

}

func TestGetCategoryFailed(t *testing.T) {

}

func TestDeleteCategorySucces(t *testing.T) {

}
func TestDeleteCategoryFailed(t *testing.T) {

}

func TestCategoriesSucces(t *testing.T) {

}

func TestUnauthorized(t *testing.T) {

}
