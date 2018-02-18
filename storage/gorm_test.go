package storage

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	"github.com/koschos/gols/domain"
	"github.com/stretchr/testify/assert"
)

func TestSaveAndFind(t *testing.T) {
	var err error

	link := domain.LinkModel{"slug1", "url1", "urlhash1"}

	db := createDb()
	repository := GormLinkRepository{*db}

	// test Save
	err = repository.Save(&link)
	assert.Nil(t, err, "Save() returns error")

	// test Find
	link1, err := repository.Find("slug1")
	assert.Nil(t, err, "Find() returns error")
	assert.Equal(t, "slug1", link1.Slug, "not found by slug 'slug1'")

	// test FindByUrlHash
	link2, err := repository.FindByUrlHash("urlhash1")
	assert.Nil(t, err, "FindByUrlHash() returns error")
	assert.Equal(t, "urlhash1", link2.UrlHash, "not found by url_hash 'urlhash1'")
}
