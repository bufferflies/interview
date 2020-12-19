package hash

import (
	"hash"

	"github.com/spaolacci/murmur3"
)

type Hash struct {
	size uint64
	hash hash.Hash64
}

// size 4==>16  8==>256
func NewHash(size int) *Hash {
	return &Hash{
		hash: murmur3.New64(),
		size: uint64(1<<size - 1),
	}
}
func (h *Hash) Hash(key []byte) (sum uint64, index int) {
	h.hash.Write(key)
	sum = h.hash.Sum64()
	index = int(sum & h.size)
	h.hash.Reset()
	return sum, index
}
