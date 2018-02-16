package generators

import (
	"hash"
	"encoding/hex"
)

type Md5HashGenerator struct {
	Hasher hash.Hash
}

func (g *Md5HashGenerator) GenerateHash(str string) string {
	g.Hasher.Write([]byte(str))

	return hex.EncodeToString(g.Hasher.Sum(nil))
}
