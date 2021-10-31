package main

// Difficulty:
// Hard

// Tags:
// Binary Search
// Hash

func longestDupSubstring(s string) string {
	out := ""
	for l, r := 1, len(s); l < r; {
		m := l + (r-l)>>1
		h := map[string]struct{}{}
		yes := false
		for i := 0; i <= len(s)-m; i++ {
			sub := s[i : i+m]
			if _, ok := h[sub]; ok {
				if m > len(out) {
					out = sub
				}
				yes = true
				break
			}
			h[sub] = struct{}{}
		}
		if yes {
			l = m + 1
		} else {
			r = m
		}
	}
	return out
}
