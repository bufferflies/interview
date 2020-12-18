package main

type Bloom interface{
	Filter(key []byte)bool
	AddFilter(key []byte)
}
