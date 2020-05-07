package leetcode

/*
  给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
*/

// Sliding Window
func minWindow(s string, t string) string {
	have, need := [128]int{}, [128]int{}
	for _, c := range t {
		need[c]++
	}
	var out string
	var l, r, count int
	for r < len(s) {
		if have[s[r]] < need[s[r]] {
			count++
		}
		have[s[r]]++
		for l <= r && have[s[l]] > need[s[l]] {
			have[s[l]]--
			l++
		}
		if count == len(t) && (len(out) == 0 || r+1-l < len(out)) {
			out = s[l : r+1]
		}
		r++
	}
	return out
}
