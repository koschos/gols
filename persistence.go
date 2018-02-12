package main

import "github.com/jinzhu/gorm"

type OrmLinkRepository struct {
	db gorm.DB
}

func (r *OrmLinkRepository) save(link *linkModel) {
	r.db.Save(&link)
}

func (r *OrmLinkRepository) find(link *linkModel, slug string) {
	r.db.First(&link, slug)
}
