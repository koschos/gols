package app

import (
	"github.com/gin-gonic/gin"
	"github.com/koschos/gols/domain"
	"net/http"
	"fmt"
)

type HashGeneratorInterface interface {
	GenerateHash(str string) string
}

type SlugGeneratorInterface interface {
	GenerateSlug() string
}

type App struct {
	LinkRepository domain.LinkRepositoryInterface
	SlugGenerator  SlugGeneratorInterface
	HashGenerator  HashGeneratorInterface
}

// create short link
func (app *App) CreateLink(c *gin.Context) {
	var link domain.LinkModel
	var createLink createLinkResource

	c.BindJSON(&createLink)

	urlHash := app.HashGenerator.GenerateHash(createLink.Url)
	app.LinkRepository.FindByUrlHash(&link, urlHash)

	if link.Slug != "" {
		resource := linkResource{Slug: link.Slug, Url: link.Url, UrlHash: link.UrlHash}

		c.JSON(http.StatusAlreadyReported, gin.H{
			"status": http.StatusAlreadyReported,
			"data": resource,
		})

		return
	}

	link.Slug = app.SlugGenerator.GenerateSlug()
	link.Url = createLink.Url
	link.UrlHash = urlHash

	app.LinkRepository.Save(&link)

	resource := linkResource{Slug: link.Slug, Url: link.Url, UrlHash: link.UrlHash}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": resource,
	})
}

// fetch a single short link
func (app *App) FetchLink(c *gin.Context) {
	var link domain.LinkModel

	slug := c.Param("slug")
	app.LinkRepository.Find(&link, slug)

	if link.Slug == "" {
		c.JSON(http.StatusNotFound, gin.H{
			"status": http.StatusNotFound,
			"message": fmt.Sprintf("No link found %s!", slug),
		})

		return
	}

	resource := linkResource{Slug: link.Slug, Url: link.Url, UrlHash: link.UrlHash}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data": resource,
	})
}
