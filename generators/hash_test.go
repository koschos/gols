package generators

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"crypto/md5"
)

func TestMd5HashGenerated(t *testing.T) {
	g := Md5HashGenerator{md5.New()}

	cases := make(map[string]string)
	cases["test string"] = "6f8db599de986fab7a21625b7916589c"

	for str, hash := range cases {
		assert.Equal(t, hash, g.GenerateHash(str))
	}
}
