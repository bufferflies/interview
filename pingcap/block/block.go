package block

import (
	"os"

	"git.code.oa.com/geeker/awesome-work/pingcap/common"
	"k8s.io/klog"
)

const KeySizeLength = 4

// Block
type Block struct {
	KeySize   int
	Key       []byte
	ValueSize int
	Value     []byte
	Offset    int64
	Length    int64
	Cached    bool
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
	keySize := common.BytesToInt(body[:KeySizeLength])
	key := body[KeySizeLength : KeySizeLength+keySize]
	valueSize := common.BytesToInt(body[KeySizeLength+keySize : KeySizeLength+keySize+KeySizeLength])
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

func GenerateBlock(size int, path string) {
	f, err := os.Create(path)
	if err != nil {
		klog.Errorf("create file failed,err:%v", err)
	}
	defer f.Close()
	for i := 0; i < size; i++ {
		k := common.IntToBytes(i)
		b := NewBlock(k, k)
		f.Write(Encode(b))
	}
}
