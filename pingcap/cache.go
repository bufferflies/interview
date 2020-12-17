package main

type Cache interface {
	Get(key []byte) []byte
}
