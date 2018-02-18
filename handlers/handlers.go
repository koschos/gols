package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/koschos/gols/domain"
	"github.com/koschos/gols/generators"
)

func RedirectHandler(repository domain.LinkRepositoryInterface) gin.HandlerFunc {
	handleFunc := func(c *gin.Context) {

		slug := c.Param("slug")

		link, err := repository.Find(slug)

		if err != nil {
			c.String(http.StatusInternalServerError, "Find error")

			return
		}

		if link.Slug == "" {
			c.String(http.StatusNotFound, "Not found")

			return
		}

		c.Redirect(http.StatusMovedPermanently, link.Url)
	}

	return gin.HandlerFunc(handleFunc)
}

func CreateLinkHandler(hashGenerator generators.HashGeneratorInterface, slugGenerator generators.SlugGeneratorInterface, repository domain.LinkRepositoryInterface) gin.HandlerFunc {
	handleFunc := func(c *gin.Context) {

		var createLink CreateLinkResource

		c.BindJSON(&createLink)

		urlHash := hashGenerator.GenerateHash(createLink.Url)
		link, err := repository.FindByUrlHash(urlHash)

		if err != nil {
			c.String(http.StatusInternalServerError, "FindByUrlHash error")

			return
		}

		if link.Slug != "" {
			c.JSON(http.StatusAlreadyReported, gin.H{
				"status": http.StatusAlreadyReported,
				"data":   GetLinkResource{Slug: link.Slug, Url: link.Url, UrlHash: link.UrlHash},
			})

			return
		}

		link.Slug = slugGenerator.GenerateSlug()
		link.Url = createLink.Url
		link.UrlHash = urlHash

		err = repository.Create(link)

		if err != nil {
			c.String(http.StatusInternalServerError, "Create error")

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status": http.StatusCreated,
			"data":   GetLinkResource{Slug: link.Slug, Url: link.Url, UrlHash: link.UrlHash},
		})
	}

	return gin.HandlerFunc(handleFunc)
}
