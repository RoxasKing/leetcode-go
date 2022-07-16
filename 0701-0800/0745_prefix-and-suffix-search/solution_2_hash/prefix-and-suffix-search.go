package main

// Difficulty:
// Hard

// Tags:
// Hash

type WordFilter struct {
	h map[string]map[string]int
}

func Constructor(words []string) WordFilter {
	h := map[string]map[string]int{}
	for i, word := range words {
		n := len(word)
		for l := 1; l <= n; l++ {
			pref := word[:l]
			if _, ok := h[pref]; !ok {
				h[pref] = map[string]int{}
			}
			for r := n - 1; r >= 0; r-- {
				suff := word[r:]
				h[pref][suff] = i
			}
		}
	}
	return WordFilter{h}
}

func (this *WordFilter) F(pref string, suff string) int {
	if _, ok := this.h[pref]; !ok {
		return -1
	}
	if _, ok := this.h[pref][suff]; !ok {
		return -1
	}
	return this.h[pref][suff]
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
