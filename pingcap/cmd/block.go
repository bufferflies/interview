package cmd

import (
	"git.code.oa.com/geeker/awesome-work/pingcap/common"
)

// Block
type Block struct {
	KeySize   int
	Key       []byte
	ValueSize int
	Value     []byte
}

// NewBlock
func NewBlock(key []byte, value []byte) Block {
	return Block{
		KeySize:   len(key),
		Key:       key,
		ValueSize: len(value),
		Value:     value,
	}
}

// Encode
func Decode(body []byte) Block {
	keySize := common.BytesToInt(body[:4])
	key := body[4 : 4+keySize]
	valueSize := common.BytesToInt(body[4+keySize : 4+keySize+4])
	value := body[4+keySize+4:]
	return Block{
		KeySize:   keySize,
		Key:       key,
		ValueSize: valueSize,
		Value:     value,
	}
}

// Decode
func Encode(b Block) []byte {
	length := b.ValueSize + b.KeySize + 4*2
	ret := make([]byte, length)
	copy(ret[0:4], common.IntToBytes(b.KeySize))
	copy(ret[4:4+b.KeySize], b.Key)
	copy(ret[4+b.KeySize:4+b.KeySize+4], common.IntToBytes(b.ValueSize))
	copy(ret[4+b.KeySize+4:], b.Value)
	return ret
}
