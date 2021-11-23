package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Bit Maniputlation
// Sorting

func maxProduct(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) > len(words[j]) })
	n := len(words)
	mask := make([]int, n)
	for i, word := range words {
		for j := range word {
			mask[i] |= 1 << int(word[j]-'a')
		}
	}
	out := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && len(words[i])*len(words[j]) > out; j++ {
			if mask[i]^mask[j] == mask[i]+mask[j] {
				out = len(words[i]) * len(words[j])
			}
		}
	}
	return out
}
