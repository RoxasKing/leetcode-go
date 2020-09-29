package main

/*
  给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
*/

// Sliding Window
func minWindow(s string, t string) string {
	out := s
	count := len(t)
	save := [128]int{}
	for i := range t {
		save[t[i]]++
	}
	for l, r := 0, 0; r < len(s); r++ {
		if save[s[r]] > 0 {
			count--
		}
		save[s[r]]--
		for l < r && save[s[l]] < 0 {
			save[s[l]]++
			l++
		}
		if count == 0 && r+1-l < len(out) {
			out = s[l : r+1]
		}
	}
	if count > 0 {
		return ""
	}
	return out
}
