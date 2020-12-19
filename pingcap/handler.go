package main

import (
	"bytes"
	"os"
	"strconv"

	"git.code.oa.com/geeker/awesome-work/pingcap/domain"
	"git.code.oa.com/geeker/awesome-work/pingcap/skipList"
	"k8s.io/klog"
)

type WriteLoop struct {
	id         int
	blockChan  chan domain.Entry
	stopCh     chan struct{}
	bufferSize int

	list  skipList.SkipList
	count int
	path  string
}

func NewWriteLoop(size int, id int, level int, stop chan struct{}, path string) *WriteLoop {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		klog.Fatalf("create dir failed ,err%v", err)
	}
	return &WriteLoop{
		blockChan:  make(chan domain.Entry, size),
		stopCh:     stop,
		id:         id,
		list:       skipList.Constructor(level),
		bufferSize: 1 << level,
		path:       path,
	}
}
func (w *WriteLoop) start() {
	for {
		select {
		case m := <-w.blockChan:
			if w.list.Size() >= w.bufferSize {
				w.write()
			}
			w.list.Add(m.Hash, &m)
		case <-w.stopCh:
			klog.Infof("write %d stop", w.id)
			w.write()
			return
		}
	}
}
func (w *WriteLoop) Send(entry domain.Entry) {
	w.blockChan <- entry
}
func (w *WriteLoop) write() {
	file := w.path + "/segment_" + strconv.Itoa(w.id) + "_" + strconv.Itoa(w.count) + ".index"
	f, err := os.Create(file)
	if err != nil {
		klog.Errorf("create file error:%v", err)
	}
	defer f.Close()
	b := bytes.Buffer{}
	for _, v := range w.list.ToValues() {
		_, err := b.Write(v.Encode())
		if err != nil {
			klog.Errorf("write entry failed err:%v", err)
		}
	}
	f.Write(b.Bytes())
	w.count++
	w.list = skipList.Constructor(w.bufferSize)
}
