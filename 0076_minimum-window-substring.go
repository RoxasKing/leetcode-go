package leetcode

/*
  给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。
*/

func minWindow(s string, t string) string {
	have, need := [128]int{}, [128]int{}
	for i := range t {
		need[t[i]]++
	}
	var out string
	for i, j, count := 0, 0, 0; j < len(s); j++ {
		if have[s[j]] < need[s[j]] {
			count++
		}
		have[s[j]]++
		for i <= j && have[s[i]] > need[s[i]] {
			have[s[i]]--
			i++
		}
		if count == len(t) && (len(out) == 0 || j+1-i < len(out)) {
			out = s[i : j+1]
		}
	}
	return out
}
