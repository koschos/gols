package generators

import (
	"time"
	"math/rand"
)

// Slug generator
type RandomSlugGenerator struct {
	Length  int
	Charset string
}

func (g *RandomSlugGenerator) GenerateSlug() string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	charset := g.createShuffledCharset(seededRand)

	b := make([]byte, g.Length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return string(b)
}

func (g *RandomSlugGenerator) createShuffledCharset(r *rand.Rand) string {
	runes := []rune(g.Charset)
	N := len(runes)
	for i := 0; i < N; i++ {
		// choose index uniformly in [i, N-1]
		r := i + r.Intn(N-i)
		runes[r], runes[i] = runes[i], runes[r]
	}

	return string(runes)
}
