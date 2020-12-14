package _5

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	s := "babad"
	assert.Equal(t, "bab", longestPalindrome(s))
	assert.Equal(t, "baab", longestPalindrome("baab"))
	assert.Equal(t, "bb", longestPalindrome("cbbd"))
	assert.Equal(t, "bb", longestPalindrome("bb"))
}
func longestPalindrome(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		odd := expand(s, i, i+1)
		nodd := expand(s, i, i)
		if odd*2 > len(result) {
			result = s[i-odd+1 : i+1+odd]
		}
		if nodd*2-1 > len(result) {
			result = s[i-nodd+1 : i+nodd]
		}
	}

	return result
}

func expand(s string, left int, right int) int {
	length := 0
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			return length
		}
		left--
		right++
		length++
	}
	return length
}
