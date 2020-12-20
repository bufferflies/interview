package main

import (
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"git.code.oa.com/geeker/awesome-work/pingcap/domain"
)

// 10 entry  each segment has less than 8 entry
// first segment should be
// 0 -1
// 0,0
//... 6,6
// second segment should be
// 7 7
// 8 ,8

func TestUnit_Write(t *testing.T) {
	te := assert.New(t)
	path := "./testdata/db"
	wg := sync.WaitGroup{}
	wg.Add(1)
	w := NewWriteLoop(10, 0, 3, path, &wg)
	go w.start()
	w.Send(domain.NewEntry(uint64(0), int64(-1)))
	for i := 0; i < 9; i++ {
		v := domain.NewEntry(uint64(i), int64(i))
		w.Send(v)
	}
	w.Close()
	wg.Wait()
	fs, err := ioutil.ReadDir(path + "/0")
	te.Nil(err)
	te.Equal(4, len(fs))
	err = os.RemoveAll(path)
	te.Nil(err)
}
