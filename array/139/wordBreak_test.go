package _39

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_word(t *testing.T) {
	s := "leetcode"
	words := []string{"leet", "code"}
	assert.True(t, wordBreak(s, words))
}
func TestUnit_apple(t *testing.T) {
	s := "applepenapplee"
	words := []string{"apple", "pen"}
	assert.False(t, wordBreak(s, words))
}
func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool, len(wordDict))
	for _, v := range wordDict {
		dict[v] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true

	// j 为切割点 只要分割在word单词中
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dict[s[j:i]] && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
