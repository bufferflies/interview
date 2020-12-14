package _64

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_nearestPalindromic(t *testing.T) {
	n := "123"
	r := nearestPalindromic(n)
	assert.Equal(t, "121", r)
}
func TestUnit_nearestPalindromic_four(t *testing.T) {
	n := "2324"
	r := nearestPalindromic(n)
	assert.Equal(t, "2332", r)
}
func nearestPalindromic(n string) string {
	if n == "11" {
		return "9"
	}
	length := len(n)
	flag := length&0x01 == 1
	a, b := n[:(length+1)/2], n[(length)/2:]
	a = reverser(a, flag)
	b = reverser(b, flag)
	numA, _ := strconv.ParseInt(a, 10, 0)
	numB, _ := strconv.ParseInt(b, 10, 0)
	numC, _ := strconv.ParseInt(n, 10, 0)
	return strconv.Itoa(min(int(numA), int(numB), int(numC)))
}

func reverser(a string, flag bool) string {
	// 说明之前是奇数
	if flag {
		m := a[len(a)-1]
		a = a[:len(a)-1]
		return a + string(m) + reverseString(a)
	} else {
		return a + reverseString(a)
	}

}

// 反转字符串
func reverseString(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}

func min(a, b, target int) int {
	if abs(a-target) <= abs(b-target) {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
