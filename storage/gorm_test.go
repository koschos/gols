package storage

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"github.com/koschos/gols/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveAndFind(t *testing.T) {

	link := domain.LinkModel{"slug1", "url1", "urlhash1"}

	db := createDb()
	repository := GormLinkRepository{*db}

	repository.Save(&link)

	link = domain.LinkModel{}
	repository.Find(&link, "slug1")

	assert.Equal(t, "slug1", link.Slug, "not found by slug 'slug1'")

	link = domain.LinkModel{}
	repository.FindByUrlHash(&link, "urlhash1")

	assert.Equal(t, "urlhash1", link.UrlHash, "not found by url_hash 'urlhash1'")
}
