package hash

import (
	"hash"

	"github.com/spaolacci/murmur3"
)

type Hash struct {
	work uint64
	hash hash.Hash64
}

// size 4==>16  8==>256
func NewHash(work int) *Hash {
	//n := int(math.Log2(float64(work)))
	return &Hash{
		hash: murmur3.New64(),
		work: uint64(work - 1),
	}
}
func (h *Hash) Hash(key []byte) (sum uint64, index int) {
	h.hash.Write(key)
	sum = h.hash.Sum64()
	index = int(sum & h.work)
	h.hash.Reset()
	return sum, index
}
