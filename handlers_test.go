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

// Test that a GET request to the home page returns the home page with
// the HTTP code 200 for an unauthenticated user
func TestPingHandler(t *testing.T) {
	r := gin.Default()

	r.GET("/ping", pingHandler)

	// Create a request to send to the above route
	req, _ := http.NewRequest("GET", "/ping", nil)

	// Create a response recorder
	w := httptest.NewRecorder()

	// Create the service and process the above request.
	r.ServeHTTP(w, req)

	// Test that the http status code is 200
	if w.Code != http.StatusOK {
		t.Error("status check failed")
	}

	expected := `{"status":"OK"}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)
}

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

	testApp := &App{&InMemoryRepository{[]linkModel{}}}

	r.POST("/", testApp.createLink)

	body := strings.NewReader(`{"url":"http://test.com"}`)

	req, _ := http.NewRequest("POST", "/", body)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	expected := `{"status":201,"data":{"slug":"slug","url":"http://test.com","url_hash":"hash"}}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)
}