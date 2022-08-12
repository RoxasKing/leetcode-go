package main

import "sort"

// Difficulty:
// Easy

// Tags:
// Sorting
// Hash

func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) > len(words[j]) })
	h := map[string]struct{}{}
	o := []string{}
	for _, w := range words {
		if _, ok := h[w]; ok {
			o = append(o, w)
		}
		for k := 1; k < len(w); k++ {
			for i := 0; i+k <= len(w); i++ {
				h[w[i:i+k]] = struct{}{}
			}
		}
	}
	return o
}
