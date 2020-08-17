package main

/*
  给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
*/

// Sliding Window
func lengthOfLongestSubstring(s string) int {
	count := [128]int{}
	var max int
	var l, r int
	for r < len(s) {
		count[s[r]]++
		for l <= r && count[s[r]] > 1 {
			count[s[l]]--
			l++
		}
		max = Max(max, r+1-l)
		r++
	}
	return max
}

func lengthOfLongestSubstring2(s string) int {
	var save [128]int
	var max int
	l, r := 0, 1
	for max < len(s)-l && r <= len(s) {
		l = Max(l, save[s[r-1]])
		max = Max(max, r-l)
		save[s[r-1]] = r
		r++
	}
	return max
}

func longestSubstring(s string) string {
	var save [128]int
	var str string
	for l, r := 0, 0; l < len(s)-len(str) && r < len(s); r++ {
		l = Max(l, save[s[r]])
		if r+1-l > len(str) {
			str = s[l : r+1]
		}
		save[s[r]] = r + 1
	}
	return str
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
