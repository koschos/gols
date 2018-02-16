package app

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"strings"
	"github.com/koschos/gols/mocks"
	"github.com/koschos/gols/domain"
)

func TestFetchLink(t *testing.T) {
	r := gin.Default()

	var linkList = []domain.LinkModel{
		{Slug:"slug1", Url:"http://url1.com", UrlHash:"urlhash1"},
	}

	testApp := &App{
		&mocks.InMemoryRepository{linkList},
		&mocks.MockSlugGenerator{},
		&mocks.MockHashGenerator{},
	}

	r.GET("/:slug", testApp.FetchLink)

	req, _ := http.NewRequest("GET", "/slug1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"status":200,"data":{"slug":"slug1","url":"http://url1.com","url_hash":"urlhash1"}}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)
}

func TestCreateLink(t *testing.T) {
	r := gin.Default()

	repository := &mocks.InMemoryRepository{[]domain.LinkModel{}}
	testApp := &App{
		repository,
		&mocks.MockSlugGenerator{"slug2"},
		&mocks.MockHashGenerator{"urlhash2"},
	}

	assert.Len(t, repository.Links, 0)

	r.POST("/", testApp.CreateLink)

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

	testApp := &App{
		repository,
		&mocks.MockSlugGenerator{"rand_slug2"},
		&mocks.MockHashGenerator{"urlhash1"},
	}

	assert.Len(t, repository.Links, 1)

	r.POST("/", testApp.CreateLink)

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