package cache

type Cache interface {
	Get(key []byte) (value []byte, err error)
}
