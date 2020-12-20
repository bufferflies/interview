package main

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_Server_Pre(t *testing.T) {
	te := assert.New(t)
	c := Config{
		Size:   100,
		Level:  2,
		Path:   "./testdata/db",
		Worker: 1,
		Src:    "./testdata/test.log",
	}
	s := NewServer(c)
	s.PreHandler()
	fs, err := ioutil.ReadDir(c.Path + "/0")
	te.Nil(err)
	te.Equal(50, len(fs))
	err = os.RemoveAll(c.Path)
	te.Nil(err)
}
