package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"status":"up"}`, w.Body.String())
}

func TestGetAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `[{"id":1,"title":"Blue Train","artist":"John Coltrane","price":56.99},{"id":2,"title":"Jeru","artist":"Gerry Mulligan","price":17.99},{"id":3,"title":"Sarah Vaughan and Clifford Brown","artist":"Sarah Vaughan","price":39.99}]`, w.Body.String())
}

func TestGetAlbumByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/1", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, `{"id":1,"title":"Blue Train","artist":"John Coltrane","price":56.99}`, w.Body.String())
}

func TestGetAlbumByIDNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/4", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, `{"message":"album not found"}`, w.Body.String())
}

func TestGetAlbumByIDBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/albums/a", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"message":"strconv.Atoi: parsing \"a\": invalid syntax"}`, w.Body.String())
}

func TestPostAlbums(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"id":4,"title":"Blue Train","artist":"John Coltrane","price":56.99}`)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, `{"id":4,"title":"Blue Train","artist":"John Coltrane","price":56.99}`, w.Body.String())
}

func TestPostAlbumsBadRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	w := httptest.NewRecorder()
	var jsonStr = []byte(`{"id":4,"title":"Blue Train","artist":"John Coltrane","price":"56.99"}`)
	req, _ := http.NewRequest("POST", "/albums", bytes.NewBuffer(jsonStr))
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
	assert.Equal(t, `{"message":"json: cannot unmarshal string into Go struct field album.price of type float64"}`, w.Body.String())
}
