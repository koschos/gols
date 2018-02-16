package main

import (
	"github.com/jinzhu/gorm"
	"github.com/koschos/gols/domain"
)

// Gorm
type OrmLinkRepository struct {
	db gorm.DB
}

func (r *OrmLinkRepository) Save(link *domain.LinkModel) {
	r.db.Save(&link)
}

func (r *OrmLinkRepository) Find(link *domain.LinkModel, slug string) {
	r.db.First(&link, slug)
}

func (r *OrmLinkRepository) FindByUrlHash(link *domain.LinkModel, urlHash string) {
	r.db.First(&link, "url_hash = ?", urlHash)
}
