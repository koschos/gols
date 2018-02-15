package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

type App struct {
	linkRepository linkRepositoryInterface
}

// create short link
func (app *App) createLink(c *gin.Context) {
	var link linkModel

	c.BindJSON(&link)

	// TODO use hasher and slug generator.
	link.Slug = "slug"
	link.UrlHash = "hash"

	app.linkRepository.save(&link)

	resource := linkResource{Slug: link.Slug, Url: link.Url, UrlHash: link.UrlHash}

	c.JSON(http.StatusCreated, gin.H{
		"status": http.StatusCreated,
		"data": resource,
	})
}

// fetch a single short link
func (app *App) fetchLink(c *gin.Context) {
	var link linkModel

	slug := c.Param("slug")
	app.linkRepository.find(&link, slug)

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
