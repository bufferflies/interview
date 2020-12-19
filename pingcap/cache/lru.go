package cache

type Lru struct {
}

func (l Lru) Get(key []byte) (value []byte, err error) {
	panic("implement me")
}
