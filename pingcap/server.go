package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"

	"git.code.oa.com/geeker/awesome-work/pingcap/cache"

	"git.code.oa.com/geeker/awesome-work/pingcap/constant"

	"git.code.oa.com/geeker/awesome-work/pingcap/skipList"

	"git.code.oa.com/geeker/awesome-work/pingcap/block"
	"git.code.oa.com/geeker/awesome-work/pingcap/domain"
	"git.code.oa.com/geeker/awesome-work/pingcap/hash"
	"git.code.oa.com/geeker/awesome-work/pingcap/storage"
	"k8s.io/klog"
)

// (size int, id int, level int, stop chan struct{}, path string)
type Config struct {
	// chan size
	Size int
	// buffer size:2^Level default is 64M(16)
	Level int
	// segment location
	Path string
	// worker
	Worker int
	// Src
	Src       string
	Cache     string
	CacheSize int
}

type Server struct {
	config   Config
	register RegisterHandler
	read     *storage.Read
	Hash     *hash.Hash
	wg       *sync.WaitGroup
	cache    cache.Cache
}
type RegisterHandler struct {
	handlerMap map[int]*WriteLoop
}

func NewServer(config Config) Server {
	r := storage.NewRead(config.Src)
	if r == nil {
		klog.Fatal("open src file  failed, please check ")
	}

	s := Server{
		config: config,
		register: RegisterHandler{
			handlerMap: make(map[int]*WriteLoop),
		},
		read: r,
		Hash: hash.NewHash(config.Worker),
		wg:   &sync.WaitGroup{},
	}
	s.wg.Add(config.Worker)
	for i := 0; i < config.Worker; i++ {
		w := NewWriteLoop(config.Size, i, config.Level, config.Path, s.wg)
		s.register.handlerMap[i] = w
	}

	if config.Cache != "" {
		s.cache = cache.Constructor(config.CacheSize, config.Cache)
		if s.cache == nil {
			klog.Errorf("cache is error ,not support:%s", config.Cache)
		}
	}
	return s
}
func (s *Server) Start() {
	s.PreHandler()
}
func (s *Server) PreHandler() {
	klog.Infof("server pre handler start ")
	for _, v := range s.register.handlerMap {
		go v.start()
	}
	b, err := s.read.ReadNext(0)
	for offset := int64(0); err == nil; {
		hash, id := s.Hash.Hash(b.Key)
		entry := convert(b, hash)
		s.send(id, entry)
		offset = b.Offset + b.Length
		b, err = s.read.ReadNext(offset)
	}
	s.close()
	s.wg.Wait()
	klog.Infof("server pre handler finish ")
}

func (s *Server) FindBlock(key []byte) (b *block.Block, err error) {
	if v := s.cache.Get(key); v != nil {
		b := block.Decode(v)
		b.Cached = true
		return &b, nil
	}
	indexes, err := s.findIndex(key)
	if err != nil {
		return nil, err
	}
	for _, entries := range indexes {
		for _, v := range entries.Entries {
			t, err := s.read.ReadNext(v.Offset)
			if err != nil {
				return nil, err
			}
			if bytes.Equal(t.Key, key) {
				s.cache.Add(key, block.Encode(t))
				return &t, nil
			}
		}
	}
	return nil, constant.KeyIsNotExistError
}

func (s *Server) findIndex(key []byte) ([]*domain.Index, error) {

	hash, _ := s.Hash.Hash(key)
	indexFiles, err := s.findBlooms(key)
	if err != nil {
		return nil, err
	}
	if len(indexFiles) == 0 {
		return nil, constant.KeyIsNotExistError
	}
	ret := make([]*domain.Index, 0)
	for i := 0; i < len(indexFiles); i++ {
		index := &domain.Index{}
		body, err := storage.Search(indexFiles[i], hash)
		if err != nil {
			return nil, err
		}
		index.Decode(body)
		ret = append(ret, index)
	}
	return ret, nil
}

func (s *Server) findBlooms(key []byte) ([]string, error) {
	_, id := s.Hash.Hash(key)
	files, err := ioutil.ReadDir(s.config.Path + "/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	bloomFiles := make([]os.FileInfo, len(files)>>1)
	for i, j := 0, 0; i < len(files); i++ {
		if strings.HasSuffix(files[i].Name(), "bloom") {
			bloomFiles[j] = files[i]
			j++
		}
	}
	ret := make([]string, 0)
	for i := 0; i < len(bloomFiles); i++ {
		location := s.GetDbFile(id, bloomFiles[i].Name())
		slot := skipList.FromFile(location)
		if slot.Test(key) {
			ret = append(ret, slot.File)
		}
	}
	return ret, nil
}
func (s *Server) GetDbFile(id int, name string) string {
	return s.config.Path + "/" + strconv.Itoa(id) + "/" + name
}

func (s *Server) send(id int, entry domain.Entry) {
	s.register.handlerMap[id].Send(entry)
}
func convert(b block.Block, hash uint64) domain.Entry {
	return domain.Entry{
		Offset: b.Offset,
		Key:    b.Key,
		Hash:   hash,
	}
}

func (s *Server) close() {
	for _, v := range s.register.handlerMap {
		v.Close()
	}
	//s.read.Close()
}
