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
			prefix := word[:l]
			if _, ok := h[prefix]; !ok {
				h[prefix] = map[string]int{}
			}
			for r := n - 1; r >= 0; r-- {
				suffix := word[r:]
				h[prefix][suffix] = i
			}
		}
	}
	return WordFilter{h}
}

func (this *WordFilter) F(prefix string, suffix string) int {
	if _, ok := this.h[prefix]; !ok {
		return -1
	}
	if _, ok := this.h[prefix][suffix]; !ok {
		return -1
	}
	return this.h[prefix][suffix]
}

/**
 * Your WordFilter object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.F(prefix,suffix);
 */
