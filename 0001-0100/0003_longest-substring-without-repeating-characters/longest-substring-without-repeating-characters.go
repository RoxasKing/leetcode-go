package main

// Tags:
// Sliding Window
func lengthOfLongestSubstring(s string) int {
	out := 0
	cnt := [128]int{}
	for l, r := 0, 0; r < len(s); r++ {
		cnt[s[r]]++
		for l < r && cnt[s[r]] > 1 {
			cnt[s[l]]--
			l++
		}
		if r+1-l > out {
			out = r + 1 - l
		}
	}
	return out
}
