package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestUnit_Should_Codec_Entry
func TestUnit_Should_Codec_Entry(t *testing.T) {
	te := assert.New(t)
	entries := []Entry{
		{
			hash:   uint64(1),
			offset: int64(1),
		},
		{
			hash:   uint64(1 << 10),
			offset: int64(1 << 10),
		},
	}
	for _, v := range entries {
		body := v.Encode()
		r := Entry{}
		r.Decode(body)
		te.Equal(v, r)
	}
	index := Index{Entries: entries}
	body := index.Encode()
	r := Index{}
	r.Decode(body)
	te.Equal(index, r)

}
