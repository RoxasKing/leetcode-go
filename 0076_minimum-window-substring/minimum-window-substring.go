package main

/*
  给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
*/

// Sliding Window
func minWindow(s string, t string) string {
	out := s
	var count int
	have := [128]int{}
	for i := range t {
		have[t[i]]++
	}
	for l, r := 0, 0; r < len(s); r++ {
		if have[s[r]] > 0 {
			count++
		}
		have[s[r]]--
		for l < r && have[s[l]] < 0 {
			have[s[l]]++
			l++
		}
		if count == len(t) && r+1-l < len(out) {
			out = s[l : r+1]
		}
	}
	if count < len(t) {
		return ""
	}
	return out
}
