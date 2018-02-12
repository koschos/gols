package main

// Model describes a todoModel type
type linkModel struct {
	Slug    string `gorm:"primary_key;type:varchar(6)"`
	Url     string `gorm:"type:text;not null" json:"url"`
	UrlHash string `gorm:"index:idx_url_hash;not null"`
}

func (linkModel) TableName() string {
	return "link"
}

// transformedModel represents a formatted resource
type linkResource struct {
	Slug    string `json:"slug"`
	Url     string `json:"url"`
	UrlHash string `json:"url_hash"`
}

// Repository
type linkRepositoryInterface interface {
	save(link *linkModel)
	find(link *linkModel, slug string)
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
