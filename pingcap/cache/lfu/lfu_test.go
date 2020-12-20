package lfu

import (
	"testing"

	"git.code.oa.com/geeker/awesome-work/pingcap/common"
	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	te := assert.New(t)
	lFUCache := Constructor(2)
	lFUCache.Add(common.IntToBytes(1), common.IntToBytes(1))
	lFUCache.Add(common.IntToBytes(2), common.IntToBytes(2))
	te.Equal(common.IntToBytes(1), lFUCache.Get(common.IntToBytes(1))) // 返回 1
	lFUCache.Add(common.IntToBytes(3), common.IntToBytes(3))           // 去除键 2
	te.Nil(lFUCache.Get(common.IntToBytes(2)))                         // 返回 -1（未找到）
	te.Equal(common.IntToBytes(3), lFUCache.Get(common.IntToBytes(3))) // 返回 3
	lFUCache.Add(common.IntToBytes(4), common.IntToBytes(4))           // 去除键 1
	te.Nil(lFUCache.Get(common.IntToBytes(1)))                         // 返回 -1（未找到）
	te.Equal(common.IntToBytes(3), lFUCache.Get(common.IntToBytes(3))) // 返回 3
	te.Equal(common.IntToBytes(4), lFUCache.Get(common.IntToBytes(4))) // 返回 4

}
