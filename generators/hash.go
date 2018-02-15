package generators

import (
	"hash"
	"encoding/hex"
)

type Md5HashGenerator struct {
	hasher hash.Hash
}

func (g *Md5HashGenerator) GenerateHash(str string) string {
	g.hasher.Write([]byte(str))

	return hex.EncodeToString(g.hasher.Sum(nil))
}
