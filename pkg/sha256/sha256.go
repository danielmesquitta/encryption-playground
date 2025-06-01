package sha256

import "crypto/sha256"

type Sha256 struct{}

func NewSha256() *Sha256 {
	return &Sha256{}
}

func (s *Sha256) Hash(data []byte) []byte {
	sum := sha256.Sum256(data)
	return sum[:]
}
