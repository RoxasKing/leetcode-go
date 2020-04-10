package leetcode

import "strings"

/*
  给定一个字符串，逐个翻转字符串中的每个单词。
*/

func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	r := len(s) - 1
	var words []string
	for l := len(s) - 1; l >= 0; l-- {
		if s[l] == ' ' {
			words = append(words, s[l+1:r+1])
			for l >= 0 && s[l] == ' ' {
				l--
			}
			r = l
		}
	}
	words = append(words, s[:r+1])
	return strings.Join(words, " ")
}
