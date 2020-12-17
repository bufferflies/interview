package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Decode_Should_Encode(t *testing.T) {
	te := assert.New(t)
	b := NewBlock([]byte("test"), []byte("fake"))
	body := Encode(b)
	r := Decode(body)
	te.Equal(b, r)
}
