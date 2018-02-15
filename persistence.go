package main

import "github.com/jinzhu/gorm"

// Gorm
type OrmLinkRepository struct {
	db gorm.DB
}

func (r *OrmLinkRepository) save(link *linkModel) {
	r.db.Save(&link)
}

func (r *OrmLinkRepository) find(link *linkModel, slug string) {
	r.db.First(&link, slug)
}

// In memory for testing
type InMemoryRepository struct {
	links []linkModel
}

func (r *InMemoryRepository) save(link *linkModel) {
	r.links = append(r.links, *link)
}

func (r *InMemoryRepository) find(link *linkModel, slug string) {
	for _, l := range r.links {
		if l.Slug == slug {
			link.Slug = l.Slug
			link.Url = l.Url
			link.UrlHash = l.UrlHash
		}
	}
}
