package main

// Difficulty:
// Medium

// Tags:
// Hash

func longestPalindrome(words []string) int {
	h := map[string]int{}
	o, x := 0, 0
	for _, w := range words {
		h[w]++
	}
	for w := range h {
		ok := true
		for l, r := 0, len(w)-1; l < r; l, r = l+1, r-1 {
			if w[l] != w[r] {
				ok = false
				break
			}
		}
		a := []rune(w)
		for l, r := 0, len(w)-1; l < r; l, r = l+1, r-1 {
			a[l], a[r] = a[r], a[l]
		}
		p, q := h[w], h[string(a)]
		if ok {
			o += p * len(w)
			if p&1 == 1 {
				o -= len(w)
				x = max(x, len(w))
			}
			continue
		}
		o += min(p, q) * len(w)
	}
	return o + x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
