package cache

import (
	"git.code.oa.com/geeker/awesome-work/pingcap/cache/lfu"
	"git.code.oa.com/geeker/awesome-work/pingcap/cache/lru"
)

type Cache interface {
	Get(key []byte) (value []byte)
	Add(key []byte, value []byte)
}

func Constructor(capacity int, t string) Cache {
	switch t {
	case "lru":
		ret := lru.Constructor(capacity)
		return &ret
	case "lfu":
		ret := lfu.Constructor(capacity)
		return &ret
	default:
		return nil
	}
}
