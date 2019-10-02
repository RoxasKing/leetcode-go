package _001_0025

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	dictionary := make([]int, 128)
	res := 0
	for i, j := 0, 0; i < len(s) && j < len(s); j++ {
		p := s[j]
		if dictionary[p] > i {
			i = dictionary[p]
		}
		num := j - i + 1
		if num > res {
			res = num
		}
		dictionary[p] = j + 1
	}
	return res
}

func TestLengthOfLongestSubString(t *testing.T) {
	s1 := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s1))
	s2 := "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s2))
	s3 := "abcdefg"
	fmt.Println(lengthOfLongestSubstring(s3))
}
