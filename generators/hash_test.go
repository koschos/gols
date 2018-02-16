package generators

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"fmt"
)

type testCase struct {
	Str  string
	Hash string
}

func TestStateless(t *testing.T) {
	g := Md5HashGenerator{}

	cases := []testCase{
		{"test string", "6f8db599de986fab7a21625b7916589c"},
		{"test string", "6f8db599de986fab7a21625b7916589c"},
	}

	for i, c := range cases {
		index := i+1
		assert.Equal(t, c.Hash, g.GenerateHash(c.Str), fmt.Sprintf("run %d failed", index))
	}
}
