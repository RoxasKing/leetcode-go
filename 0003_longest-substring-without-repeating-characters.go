package leetcode

func lengthOfLongestSubstring(s string) int {
	save := make([]int, 128)
	var max int
	l, r := 0, 1
	for max <= len(s)-l && r <= len(s) {
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
