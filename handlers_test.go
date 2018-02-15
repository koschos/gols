package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	//"net/url"
	"strings"
)

func TestFetchLinkHandler(t *testing.T) {
	r := gin.Default()

	var linkList = []linkModel{
		{Slug:"slug1", Url:"http://url1.com", UrlHash:"urlhash1"},
	}

	testApp := &App{
		&InMemoryRepository{linkList},
	}

	r.GET("/:slug", testApp.fetchLink)

	req, _ := http.NewRequest("GET", "/slug1", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"status":200,"data":{"slug":"slug1","url":"http://url1.com","url_hash":"urlhash1"}}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)
}

func TestCreateLinkHandler(t *testing.T) {
	r := gin.Default()

	repository := &InMemoryRepository{[]linkModel{}}
	testApp := &App{repository}

	r.POST("/", testApp.createLink)

	body := strings.NewReader(`{"url":"http://test.com"}`)

	req, _ := http.NewRequest("POST", "/", body)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	expected := `{"status":201,"data":{"slug":"slug","url":"http://test.com","url_hash":"hash"}}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)

	assert.Len(t, repository.links, 1)
}