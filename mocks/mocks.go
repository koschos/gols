package mocks

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
