package generators

import (
	"encoding/hex"
	"crypto/md5"
)

type Md5HashGenerator struct {}

func (g *Md5HashGenerator) GenerateHash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}
