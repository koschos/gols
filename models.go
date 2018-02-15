package main

// Model, annotated with gorm
type linkModel struct {
	Slug    string `gorm:"primary_key;type:varchar(6)"`
	Url     string `gorm:"type:text;not null"`
	UrlHash string `gorm:"index:idx_url_hash;not null"`
}

func (linkModel) TableName() string {
	return "link"
}

// Repository
type linkRepositoryInterface interface {
	save(link *linkModel)
	find(link *linkModel, slug string)
}
