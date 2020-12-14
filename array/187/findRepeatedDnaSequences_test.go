package _87

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit(t *testing.T) {
	s := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTTå"
	expect := []string{"AAAAACCCCC", "CCCCCAAAAA"}
	assert.Equal(t, expect, findRepeatedDnaSequences(s))
}
func TestUnit_AAA(t *testing.T) {
	s := "AAAAAAAAAAAAA"
	expect := []string{"AAAAAAAAAA"}
	assert.Equal(t, expect, findRepeatedDnaSequences(s))
}
func findRepeatedDnaSequences(s string) []string {
	L := 10
	set := make(map[string]struct{}, 0)
	dict := make(map[string]bool)
	for i := 0; i < len(s)-L; i++ {
		tmp := s[i : i+L]
		if dict[tmp] {
			set[tmp] = struct{}{}
		} else {
			dict[tmp] = true
		}
	}
	result := make([]string, 0)
	for i, _ := range set {
		result = append(result, i)
	}
	return result

}
