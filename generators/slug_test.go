package generators

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func TestRandomSlugGenerated(t *testing.T) {
	g := RandomSlugGenerator{6, charset}

	slug1 := g.GenerateSlug()
	slug2 := g.GenerateSlug()

	assert.Len(t, slug1, 6)
	assert.Len(t, slug2, 6)
	assert.NotEqual(t, slug1, slug2)
}
