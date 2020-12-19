package main

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"git.code.oa.com/geeker/awesome-work/pingcap/domain"
)

func TestUnit_Write(t *testing.T) {
	te := assert.New(t)
	stopCh := make(chan struct{})
	path := "./db"
	w := NewWriteLoop(10, 1, 3, stopCh, path)
	go w.start()
	w.Send(domain.NewEntry(uint64(0), int64(-1)))
	for i := 0; i < 9; i++ {
		v := domain.NewEntry(uint64(i), int64(i))
		w.Send(v)
	}
	time.Sleep(time.Second * 3)
	stopCh <- struct{}{}
	time.Sleep(time.Second * 3)
	fs, err := ioutil.ReadDir(path)
	te.Nil(err)
	te.Equal(2, len(fs))
	te.Equal(int64(16*8), fs[0].Size())
	te.Equal(int64(16*2), fs[1].Size())
}
