package domain

// Model, annotated with gorm
type LinkModel struct {
	Slug    string `gorm:"primary_key;type:varchar(6)"`
	Url     string `gorm:"type:text;not null"`
	UrlHash string `gorm:"index:idx_url_hash;not null"`
}

func (LinkModel) TableName() string {
	return "link"
}

// Repository
type LinkRepositoryInterface interface {
	Save(link *LinkModel)
	Find(link *LinkModel, slug string)
	FindByUrlHash(link *LinkModel, urlHash string)
}
