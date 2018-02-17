package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"github.com/koschos/gols/mocks"
	"github.com/koschos/gols/domain"
	"strings"
)

func TestRedirect(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{[]domain.LinkModel{
		{Slug:"slug1", Url:"http://test.com", UrlHash:"urlhash1"},
	}}

	r.GET("/:slug", RedirectHandler(repository))

	req, _ := http.NewRequest("GET", "/slug1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
}

func TestRedirectNotFound(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{[]domain.LinkModel{
		{Slug:"slug1", Url:"http://test.com", UrlHash:"urlhash1"},
	}}

	r.GET("/:slug", RedirectHandler(repository))

	req, _ := http.NewRequest("GET", "/slug2", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestFetchLink(t *testing.T) {
	r := gin.Default()

	var linkList = []domain.LinkModel{
		{Slug:"slug1", Url:"http://url1.com", UrlHash:"urlhash1"},
	}

	repository := &mocks.InMemoryRepository{linkList}

	r.GET("/:slug", FetchLinkHandler(repository))

	req, _ := http.NewRequest("GET", "/slug1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"status":200,"data":{"slug":"slug1","url":"http://url1.com","url_hash":"urlhash1"}}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)
}

func TestCreateNewLink(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{[]domain.LinkModel{}}
	hashGenerator := &mocks.MockHashGenerator{"urlhash2"}
	slugGenerator := &mocks.MockSlugGenerator{"slug2"}

	assert.Len(t, repository.Links, 0)

	r.POST("/", CreateLinkHandler(hashGenerator, slugGenerator, repository))

	body := strings.NewReader(`{"url":"http://test.com"}`)

	req, _ := http.NewRequest("POST", "/", body)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	expected := `{"status":201,"data":{"slug":"slug2","url":"http://test.com","url_hash":"urlhash2"}}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)

	assert.Len(t, repository.Links, 1)
}

func TestCreateAlreadyExistingLink(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{[]domain.LinkModel{
		{Slug:"rand_slug1", Url:"http://test.com", UrlHash:"urlhash1"},
	}}
	hashGenerator := &mocks.MockHashGenerator{"urlhash1"}
	slugGenerator := &mocks.MockSlugGenerator{"rand_slug2"}

	assert.Len(t, repository.Links, 1)

	r.POST("/", CreateLinkHandler(hashGenerator, slugGenerator, repository))

	body := strings.NewReader(`{"url":"http://test.com"}`)

	req, _ := http.NewRequest("POST", "/", body)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusAlreadyReported, w.Code)

	expected := `{"status":208,"data":{"slug":"rand_slug1","url":"http://test.com","url_hash":"urlhash1"}}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)

	assert.Len(t, repository.Links, 1)
}
