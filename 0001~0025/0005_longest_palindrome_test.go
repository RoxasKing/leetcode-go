package _001_0025

import (
	"fmt"
	"testing"
)

func longestPalindrome(s string) string {
	size, out := len(s), ""
	if size < 2 {
		return s
	}
	for i := 0; i < size; i++ {
		// 如果未遍历字符串长度小于等于当前最长回文子串长度
		if size-i <= len(out)/2 {
			break
		}
		// 记录连续相同字符最后的下标
		e := i
		for e+1 < size && s[e] == s[e+1] {
			e++
		}
		// 回文子串两翼长度
		j := 0
		for i-j-1 >= 0 && e+j+1 < size && s[i-j-1] == s[e+j+1] {
			j++
		}
		// 判断当前回文子串长度是否超过已记录最大值
		if e-i+1+2*j > len(out) {
			out = s[i-j : e+j+1]
		}
	}
	return out
}

func TestLongestPalindrome(t *testing.T) {
	s := "babad"
	fmt.Println(longestPalindrome(s))
}
