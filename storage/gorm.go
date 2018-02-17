package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/koschos/gols/domain"
)

// Gorm
type GormLinkRepository struct {
	Db gorm.DB
}

func (r *GormLinkRepository) Save(link *domain.LinkModel) {
	r.Db.Save(&link)
}

func (r *GormLinkRepository) Find(link *domain.LinkModel, slug string) {
	r.Db.First(&link, "slug = ?", slug)
}

func (r *GormLinkRepository) FindByUrlHash(link *domain.LinkModel, urlHash string) {
	r.Db.First(&link, "url_hash = ?", urlHash)
}
