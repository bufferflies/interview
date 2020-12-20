package main

import (
	"os"
	"testing"

	"git.code.oa.com/geeker/awesome-work/pingcap/common"

	"github.com/stretchr/testify/assert"
)

// test.log has 100 kv
// 1 worker
//
func TestUnit_PreHandler(t *testing.T) {
	te := assert.New(t)
	c := Config{
		Size:      10,
		Level:     2,
		Path:      "./testdata/db",
		Worker:    8,
		Src:       "./testdata/test.log",
		Cache:     "lfu",
		CacheSize: 10,
	}
	s := NewServer(c)
	s.PreHandler()

	b, err := s.FindBlock(common.IntToBytes(0))
	te.Nil(err)
	te.False(b.Cached)
	b, err = s.FindBlock(common.IntToBytes(0))
	te.Nil(err)
	te.True(b.Cached)

	_, err = s.FindBlock(common.IntToBytes(101))
	te.NotNil(err)
	_, err = s.FindBlock(common.IntToBytes(102))
	te.NotNil(err)
	os.RemoveAll(c.Path)
}

func TestUit_Read(t *testing.T) {
	te := assert.New(t)
	c := Config{
		Size:   10,
		Level:  2,
		Path:   "./testdata/db",
		Worker: 1,
		Src:    "./testdata/test.log",
	}
	s := NewServer(c)
	b, err := s.FindBlock(common.IntToBytes(0))
	te.Nil(err)
	te.Equal(b.Value, common.IntToBytes(0))

}
