package main

import (
	"git.code.oa.com/geeker/awesome-work/pingcap/block"
	"git.code.oa.com/geeker/awesome-work/pingcap/skipList"
	"k8s.io/klog"
)

type WriteLoop struct {
	id         int
	blockChan  chan block.Block
	stopCh     chan struct{}
	bufferSize int
	list       skipList.SkipList
}

func NewWriteLoop(size int, id int, bufferSize int, stop chan struct{}) *WriteLoop {
	return &WriteLoop{
		blockChan: make(chan block.Block, size),
		stopCh:    stop,
		id:        id,
		list:      skipList.Constructor(bufferSize),
	}
}
func (w *WriteLoop) start() {
	for {
		select {
		case m := <-w.blockChan:
			klog.Infof("writer %d receive msg ,key:%s,value :%s", w.id, string(m.Key), string(m.Value))
		case <-w.stopCh:
			klog.Infof("write %d stop", w.id)
		}
	}
}
func (w WriteLoop) write() {
	list := w.list
	arr := list.ToValues()

}
