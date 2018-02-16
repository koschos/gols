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
}

func (r *InMemoryRepository) Save(link *domain.LinkModel) {
	r.Links = append(r.Links, *link)
}

func (r *InMemoryRepository) Find(link *domain.LinkModel, slug string) {
	for _, l := range r.Links {
		if l.Slug == slug {
			link.Slug = l.Slug
			link.Url = l.Url
			link.UrlHash = l.UrlHash
		}
	}
}

func (r *InMemoryRepository) FindByUrlHash(link *domain.LinkModel, urlHash string) {
	for _, l := range r.Links {
		if l.UrlHash == urlHash {
			link.Slug = l.Slug
			link.Url = l.Url
			link.UrlHash = l.UrlHash
		}
	}
}
