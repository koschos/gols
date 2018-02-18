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
	"errors"
)

func TestRedirect301(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{
		Links: []domain.LinkModel{
			{Slug:"slug1", Url:"http://test.com", UrlHash:"urlhash1"},
		},
	}

	r.GET("/:slug", RedirectHandler(repository))

	req, _ := http.NewRequest("GET", "/slug1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
}

func TestRedirect404(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{
		Links: []domain.LinkModel{
			{Slug:"slug1", Url:"http://test.com", UrlHash:"urlhash1"},
		},
	}

	r.GET("/:slug", RedirectHandler(repository))

	req, _ := http.NewRequest("GET", "/slug2", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "Not found", w.Body.String())
}

func TestRedirect500r(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{
		Error: errors.New("db error"),
	}

	r.GET("/:slug", RedirectHandler(repository))

	req, _ := http.NewRequest("GET", "/slug1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateLink201(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{Links: []domain.LinkModel{}}
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

func TestCreateLink208(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{
		Links:[]domain.LinkModel{
			{Slug:"rand_slug1", Url:"http://test.com", UrlHash:"urlhash1"},
		},
	}

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

func TestCreateLink500FindError(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{
		Error: errors.New("db error"),
	}

	hashGenerator := &mocks.MockHashGenerator{"urlhash1"}
	slugGenerator := &mocks.MockSlugGenerator{"rand_slug2"}

	r.POST("/", CreateLinkHandler(hashGenerator, slugGenerator, repository))

	body := strings.NewReader(`{"url":"http://test.com"}`)

	req, _ := http.NewRequest("POST", "/", body)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "FindByUrlHash error", w.Body.String())
	assert.Len(t, repository.Links, 0)
}

func TestCreateLink500SaveError(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{
		SaveError: errors.New("db error"),
	}

	hashGenerator := &mocks.MockHashGenerator{"urlhash1"}
	slugGenerator := &mocks.MockSlugGenerator{"rand_slug2"}

	r.POST("/", CreateLinkHandler(hashGenerator, slugGenerator, repository))

	body := strings.NewReader(`{"url":"http://test.com"}`)

	req, _ := http.NewRequest("POST", "/", body)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "Save error", w.Body.String())
	assert.Len(t, repository.Links, 0)
}
