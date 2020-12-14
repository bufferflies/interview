package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Trie struct {
	isWord bool
	next   [26]*Trie
}

func Constructor() Trie {
	return Trie{}
}
func (this *Trie) Insert(word string) {
	cur := this
	for i, c := range word {
		n := c - 'a'
		if cur.next[n] == nil {
			cur.next[n] = &Trie{}
		}
		cur = cur.next[n]
		if i == len(word)-1 {
			cur.isWord = true
		}
	}
}

func (this *Trie) Search(word string) bool {
	cur := this
	for _, c := range word {
		n := c - 'a'
		if cur.next[n] == nil {
			return false
		}
		cur = cur.next[n]
	}
	return cur.isWord
}
func (this *Trie) StartsWith(prefix string) bool {
	cur := this
	for _, c := range prefix {
		n := c - 'a'
		if cur.next[n] == nil {
			return false
		}
		cur = cur.next[n]
	}
	return true
}
func TestUnit_Trie(t *testing.T) {
	te := assert.New(t)
	trie := Constructor()
	b := trie.Search("test")
	te.False(b)
	trie.Insert("test")
	b = trie.Search("tes")
	te.False(b)
	b = trie.Search("test")
	te.True(b)
}
