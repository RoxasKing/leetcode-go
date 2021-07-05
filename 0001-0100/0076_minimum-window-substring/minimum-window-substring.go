package main

// Tags:
// Sliding Window + Two Pointers
func minWindow(s string, t string) string {
	cnt := [128]int{}
	for i := range t {
		cnt[t[i]]++
	}
	out := ""
	for l, r, include := 0, 0, 0; r < len(s); r++ {
		if cnt[s[r]] > 0 {
			include++
		}
		cnt[s[r]]--
		for l <= r && cnt[s[l]] < 0 {
			cnt[s[l]]++
			l++
		}
		if include == len(t) && (out == "" || len(out) > r+1-l) {
			out = s[l : r+1]
		}
	}
	return out
}
