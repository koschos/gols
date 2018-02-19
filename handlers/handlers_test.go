package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/koschos/gols/mocks"
	"github.com/koschos/gols/domain"
	"github.com/koschos/gols/generators"
)

func TestRedirect_301(t *testing.T) {

	repo := &mocks.InMemoryRepository{
		Links: []domain.LinkModel{
			{Slug:"slug1", Url:"http://test.com", UrlHash:"urlhash1"},
		},
	}

	w := runRedirect("slug1", repo)

	assert.Equal(t, http.StatusMovedPermanently, w.Code)
}

func TestRedirect_404(t *testing.T) {

	repo := &mocks.InMemoryRepository{
		Links: []domain.LinkModel{
			{Slug:"slug1", Url:"http://test.com", UrlHash:"urlhash1"},
		},
	}

	w := runRedirect("slug2", repo)

	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Equal(t, "Not found", w.Body.String())
}

func TestRedirect_500(t *testing.T) {

	repo := &mocks.InMemoryRepository{
		Error: errors.New("db error"),
	}

	w := runRedirect("slug1", repo)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestCreateLink_400_BadRequest(t *testing.T) {

	hg := &mocks.MockHashGenerator{"urlhash1"}
	sg := &mocks.MockSlugGenerator{[]string{"slug1"}}
	repo := &mocks.InMemoryRepository{Links: []domain.LinkModel{}}

	w := runCreate(`{"wrong":"test"}`, hg, sg, repo)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Len(t, repo.Links, 0)
}

func TestCreateLink_201(t *testing.T) {

	hg := &mocks.MockHashGenerator{"urlhash2"}
	sg := &mocks.MockSlugGenerator{[]string{"slug2"}}
	repo := &mocks.InMemoryRepository{Links: []domain.LinkModel{}}

	w := runCreate(`{"url":"http://test.com"}`, hg, sg, repo)

	assert.Equal(t, http.StatusCreated, w.Code)

	expected := `{"status":201,"data":{"slug":"slug2","url":"http://test.com","url_hash":"urlhash2"}}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)

	assert.Len(t, repo.Links, 1)
}

func TestCreateLink_201_MySql1062Duplicated(t *testing.T) {

	hg := &mocks.MockHashGenerator{"urlhash1"}
	sg := &mocks.MockSlugGenerator{[]string{"rand_slug1", "rand_slug2"}}

	repo := &mocks.InMemoryRepository{
		CreateError: &mysql.MySQLError{Number: 1062},
	}

	w := runCreate(`{"url":"http://test.com"}`, hg, sg, repo)

	assert.Equal(t, http.StatusCreated, w.Code)

	expected := `{"status":201,"data":{"slug":"rand_slug2","url":"http://test.com","url_hash":"urlhash1"}}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)

	assert.Len(t, repo.Links, 1)
}

func TestCreateLink_208(t *testing.T) {

	hg := &mocks.MockHashGenerator{"urlhash1"}
	sg := &mocks.MockSlugGenerator{[]string{"rand_slug2"}}

	repo := &mocks.InMemoryRepository{
		Links:[]domain.LinkModel{
			{Slug:"rand_slug1", Url:"http://test.com", UrlHash:"urlhash1"},
		},
	}

	w := runCreate(`{"url":"http://test.com"}`, hg, sg, repo)

	assert.Equal(t, http.StatusAlreadyReported, w.Code)

	expected := `{"status":208,"data":{"slug":"rand_slug1","url":"http://test.com","url_hash":"urlhash1"}}`
	actual := w.Body.String()
	assert.JSONEq(t, expected, actual, "handler returned unexpected body: got %v want %v", expected, actual)

	assert.Len(t, repo.Links, 1)
}

func TestCreateLink_500_FindError(t *testing.T) {

	hg := &mocks.MockHashGenerator{"urlhash1"}
	sg := &mocks.MockSlugGenerator{[]string{"rand_slug2"}}

	repo := &mocks.InMemoryRepository{
		Error: errors.New("db error"),
	}

	w := runCreate(`{"url":"http://test.com"}`, hg, sg, repo)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "FindByUrlHash error", w.Body.String())
	assert.Len(t, repo.Links, 0)
}

func TestCreateLink_500_CreateError(t *testing.T) {

	hg := &mocks.MockHashGenerator{"urlhash1"}
	sg := &mocks.MockSlugGenerator{[]string{"rand_slug2"}}

	repo := &mocks.InMemoryRepository{
		CreateError: errors.New("db error"),
	}

	w := runCreate(`{"url":"http://test.com"}`, hg, sg, repo)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Equal(t, "Create error", w.Body.String())
	assert.Len(t, repo.Links, 0)
}

func runCreate(
	json string,
	hg generators.HashGeneratorInterface,
	sg generators.SlugGeneratorInterface,
	repo domain.LinkRepositoryInterface) *httptest.ResponseRecorder {

	r := gin.Default()

	r.POST("/", CreateLinkHandler(hg, sg, repo))

	body := strings.NewReader(json)

	req, _ := http.NewRequest("POST", "/", body)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}

func runRedirect(slug string, repo domain.LinkRepositoryInterface) *httptest.ResponseRecorder {
	r := gin.Default()

	r.GET("/:slug", RedirectHandler(repo))

	req, _ := http.NewRequest("GET", "/" + slug, nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	return w
}