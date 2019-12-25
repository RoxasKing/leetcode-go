package _001_0025

import (
	"fmt"
	"testing"
)

func lengthOfLongestSubstring(s string) int {
	dict := make([]int, 128)
	res := 0
	for i, j := 0, 0; i < len(s) && j < len(s); j++ {
		p := s[j]
		// 若当前字符索引大于最左边界索引，当前字符索引即为最左边界索引
		i = max(i, dict[p])
		res = max(res, j-i+1) // 判断 [i,j] 是否为最长子串
		dict[p] = j + 1       // 索引从 1 算起
	}
	return res
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func TestLengthOfLongestSubString(t *testing.T) {
	s1 := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s1))
	s2 := "bbbbb"
	fmt.Println(lengthOfLongestSubstring(s2))
	s3 := "abcdefg"
	fmt.Println(lengthOfLongestSubstring(s3))
}
