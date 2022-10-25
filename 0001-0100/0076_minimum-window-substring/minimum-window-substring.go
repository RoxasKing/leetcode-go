package main

// Difficulty:
// Hard

// Tags:
// Sliding Window

func minWindow(s string, t string) string {
	h := [128]int{}
	for i := range t {
		h[t[i]]++
	}
	o := ""
	for l, r, c := 0, 0, 0; r < len(s); r++ {
		if h[s[r]] > 0 {
			c++
		}
		h[s[r]]--
		for l <= r && h[s[l]] < 0 {
			h[s[l]]++
			l++
		}
		if c == len(t) && (o == "" || len(o) > r+1-l) {
			o = s[l : r+1]
		}
	}
	return o
}
