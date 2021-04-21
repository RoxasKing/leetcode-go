package main

/*
  给定一个字符串，你的任务是计算这个字符串中有多少个回文子串。
  具有不同开始位置或结束位置的子串，即使是由相同的字符组成，也会被视作不同的子串。

  提示：
    输入的字符串长度不会超过 1000 。
*/

func countSubstrings(s string) int {
	var out int
	for i := range s {
		out += countPalindrome(&s, i, i)
		if i < len(s)-1 && s[i+1] == s[i] {
			out += countPalindrome(&s, i, i+1)
		}
	}
	return out
}

func countPalindrome(s *string, l, r int) int {
	count := 1
	for 0 < l && r < len(*s)-1 && (*s)[l-1] == (*s)[r+1] {
		l--
		r++
		count++
	}
	return count
}
