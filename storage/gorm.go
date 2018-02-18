package storage

import (
	"github.com/jinzhu/gorm"
	"github.com/koschos/gols/domain"
)

// Gorm
type GormLinkRepository struct {
	Db gorm.DB
}

func (r *GormLinkRepository) Save(link *domain.LinkModel) (error) {
	return r.Db.Save(&link).Error
}

func (r *GormLinkRepository) Find(slug string) (*domain.LinkModel, error) {
	var link = &domain.LinkModel{}

	r.Db.First(&link, "slug = ?", slug)

	return link, r.Db.Error
}

func (r *GormLinkRepository) FindByUrlHash(urlHash string) (*domain.LinkModel, error) {
	var link = &domain.LinkModel{}

	r.Db.First(&link, "url_hash = ?", urlHash)

	return link, r.Db.Error
}
