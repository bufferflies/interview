package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Int(t *testing.T) {
	te := assert.New(t)
	var r = []struct {
		n int
	}{
		{1},
		{1 << 10},
		{1 << 20},
	}
	for _, v := range r {
		d := IntToBytes(v.n)
		e := BytesToInt(d)
		te.Equal(v.n, e)
	}
}
func TestUnit_Int64(t *testing.T) {
	te := assert.New(t)
	var r = []struct {
		n int64
	}{
		{1},
		{1 << 10},
		{1 << 20},
	}
	for _, v := range r {
		d := Int64ToBytes(v.n)
		e := BytesToInt64(d)
		te.Equal(v.n, e)
	}
}
