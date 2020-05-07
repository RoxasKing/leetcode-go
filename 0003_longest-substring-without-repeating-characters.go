package leetcode

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
	save := make([]int, 128)
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
	save := make([]int, 128)
	var str string
	l, r := 0, 1
	for len(str) <= len(s)-l && r <= len(s) {
		l = Max(l, save[s[r-1]])
		if r-l > len(str) {
			str = s[l:r]
		}
		save[s[r-1]] = r
		r++
	}
	return str
}
