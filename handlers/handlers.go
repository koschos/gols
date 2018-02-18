package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/koschos/gols/domain"
	"github.com/koschos/gols/generators"
	"github.com/go-sql-driver/mysql"
)

const duplicateEntryErrorNumber = 1062

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

func CreateLinkHandler(
	hashGenerator generators.HashGeneratorInterface,
	slugGenerator generators.SlugGeneratorInterface,
	repository domain.LinkRepositoryInterface) gin.HandlerFunc {

	handleFunc := func(c *gin.Context) {

		var createLink CreateLinkResource

		err := c.BindJSON(&createLink)

		if err != nil || createLink.Url == "" {
			c.String(http.StatusBadRequest, "Bad request")

			return
		}

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

		link.Url = createLink.Url
		link.UrlHash = urlHash

		// Create link, skip duplicated errors.
		for {

			link.Slug = slugGenerator.GenerateSlug()

			err = repository.Create(link)

			// stop on first success
			if err == nil {
				break
			}

			if isDuplicateEntryError(err) {
				continue
			}

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

func isDuplicateEntryError(err error) bool {
	mySqlError, ok := err.(*mysql.MySQLError)

	return ok && int(mySqlError.Number) == duplicateEntryErrorNumber
}