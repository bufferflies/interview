package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Hash(t *testing.T) {
	te := assert.New(t)
	var result = []struct {
		key   []byte
		hash  uint64
		index int
	}{
		{[]byte("test"), 12429135405209477533, 157},
		{[]byte("fake"), 15575225405007291005, 125},
		{[]byte("hello,world"), 1311825808656992506, 250},
	}
	h := NewHash(8)
	for _, v := range result {
		hash, index := h.Hash(v.key)
		te.Equal(v.hash, hash)
		te.Equal(v.index, index)
	}
}
