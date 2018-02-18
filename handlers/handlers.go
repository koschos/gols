package handlers

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/koschos/gols/domain"
	"github.com/koschos/gols/generators"
)

func RedirectHandler(repository domain.LinkRepositoryInterface) gin.HandlerFunc {
	handleFunc := func(c *gin.Context) {
		var slug string
		var link domain.LinkModel

		slug = c.Param("slug")

		repository.Find(&link, slug)

		if link.Slug == "" {
			c.String(http.StatusNotFound, "Not found")
		}

		c.Redirect(http.StatusMovedPermanently, link.Url)
	}

	return gin.HandlerFunc(handleFunc)
}

func FetchLinkHandler(repository domain.LinkRepositoryInterface) gin.HandlerFunc {
	handleFunc := func(c *gin.Context) {
		var link domain.LinkModel

		slug := c.Param("slug")
		repository.Find(&link, slug)

		if link.Slug == "" {
			c.JSON(http.StatusNotFound, gin.H{
				"status": http.StatusNotFound,
				"message": fmt.Sprintf("No link found %s!", slug),
			})

			return
		}

		resource := getLinkResource{Slug: link.Slug, Url: link.Url, UrlHash: link.UrlHash}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data": resource,
		})
	}

	return gin.HandlerFunc(handleFunc)
}

func CreateLinkHandler(hashGenerator generators.HashGeneratorInterface, slugGenerator generators.SlugGeneratorInterface, repository domain.LinkRepositoryInterface) gin.HandlerFunc {
	handleFunc := func(c *gin.Context) {
		var link domain.LinkModel
		var createLink createLinkResource
		var urlHash string

		c.BindJSON(&createLink)

		urlHash = hashGenerator.GenerateHash(createLink.Url)
		repository.FindByUrlHash(&link, urlHash)

		if link.Slug != "" {
			resource := getLinkResource{Slug: link.Slug, Url: link.Url, UrlHash: link.UrlHash}

			c.JSON(http.StatusAlreadyReported, gin.H{
				"status": http.StatusAlreadyReported,
				"data": resource,
			})

			return
		}

		link.Slug = slugGenerator.GenerateSlug()
		link.Url = createLink.Url
		link.UrlHash = urlHash

		repository.Save(&link)

		resource := getLinkResource{Slug: link.Slug, Url: link.Url, UrlHash: link.UrlHash}

		c.JSON(http.StatusCreated, gin.H{
			"status": http.StatusCreated,
			"data": resource,
		})
	}

	return gin.HandlerFunc(handleFunc)
}
