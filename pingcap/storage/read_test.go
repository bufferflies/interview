package storage

import (
	"testing"

	"git.code.oa.com/geeker/awesome-work/pingcap/common"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	te := assert.New(t)
	read := NewRead("./test.log")
	offset := int64(0)
	b, err := read.ReadNext(offset)
	i := 0
	for ; err == nil; i++ {
		te.Equal(common.IntToBytes(i), b.Key)
		te.Equal(common.IntToBytes(i), b.Value)
		te.Equal(len(b.Key), b.KeySize)
		te.Equal(len(b.Value), b.ValueSize)
		te.Equal(offset, b.Offset)
		offset = b.Offset + b.Length
		b, err = read.ReadNext(offset)
	}
	te.Equal(100, i)
	te.NotNil(err)
	read.Close()
}
