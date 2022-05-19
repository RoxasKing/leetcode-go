package main

// Difficulty:
// Easy

func isAlienSorted(words []string, order string) bool {
	h := [26]int{}
	for i, c := range order {
		h[c-'a'] = i
	}
	for i := 0; i < len(words)-1; i++ {
		a, b := words[i], words[i+1]
		k := min(len(a), len(b))
		if a[:k] == b[:k] {
			if len(a) > len(b) {
				return false
			}
			continue
		}
		for j := 0; j < min(len(a), len(b)); j++ {
			if h[a[j]-'a'] == h[b[j]-'a'] {
				continue
			} else if h[a[j]-'a'] < h[b[j]-'a'] {
				break
			} else {
				return false
			}
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
