package lru

import (
	"testing"

	"git.code.oa.com/geeker/awesome-work/pingcap/common"
	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	te := assert.New(t)
	cache := Constructor(2)
	cache.Add(common.IntToBytes(1), common.IntToBytes(1))
	cache.Add(common.IntToBytes(2), common.IntToBytes(2))
	r := cache.Get(common.IntToBytes(1))
	te.Equal(common.IntToBytes(1), r)
	cache.Add(common.IntToBytes(3), common.IntToBytes(3))
	r = cache.Get(common.IntToBytes(2))
	te.Nil(r)
	cache.Add(common.IntToBytes(4), common.IntToBytes(4))
	r = cache.Get(common.IntToBytes(1))
	te.Nil(r)
	r = cache.Get(common.IntToBytes(3))
	te.Equal(common.IntToBytes(3), r)
	r = cache.Get(common.IntToBytes(4))
	te.Equal(common.IntToBytes(4), r)
}
