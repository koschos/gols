package mocks

import "github.com/koschos/gols/domain"

type MockSlugGenerator struct {
	Slugs []string
}

func (g *MockSlugGenerator) GenerateSlug() string {
	slug := g.Slugs[0]
	g.Slugs = g.Slugs[1:]

	return slug
}

type MockHashGenerator struct {
	Hash string
}

func (g *MockHashGenerator) GenerateHash(str string) string {
	return g.Hash
}

// In memory for testing
type InMemoryRepository struct {
	Links       []domain.LinkModel
	Error       error
	CreateError error
}

func (r *InMemoryRepository) Create(link *domain.LinkModel) (error) {
	if r.CreateError != nil {
		err := r.CreateError
		r.CreateError = nil

		return err
	}

	r.Links = append(r.Links, *link)

	return r.CreateError
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
