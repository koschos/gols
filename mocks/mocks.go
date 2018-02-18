package mocks

import "github.com/koschos/gols/domain"

type MockSlugGenerator struct {
	Slug string
}

func (g *MockSlugGenerator) GenerateSlug() string {
	return g.Slug
}

type MockHashGenerator struct {
	Hash string
}

func (g *MockHashGenerator) GenerateHash(str string) string {
	return g.Hash
}

// In memory for testing
type InMemoryRepository struct {
	Links []domain.LinkModel
	Error error
	SaveError error
}

func (r *InMemoryRepository) Save(link *domain.LinkModel) (error) {
	if r.SaveError == nil {
		r.Links = append(r.Links, *link)
	}

	return r.SaveError
}

func (r *InMemoryRepository) Find(slug string) (*domain.LinkModel, error) {
	var link = &domain.LinkModel{}

	for _, l := range r.Links {
		if l.Slug == slug {
			link.Slug = l.Slug
			link.Url = l.Url
			link.UrlHash = l.UrlHash
		}
	}

	return link, r.Error
}

func (r *InMemoryRepository) FindByUrlHash(urlHash string) (*domain.LinkModel, error) {
	var link = &domain.LinkModel{}

	for _, l := range r.Links {
		if l.UrlHash == urlHash {
			link.Slug = l.Slug
			link.Url = l.Url
			link.UrlHash = l.UrlHash
		}
	}

	return link, r.Error
}
