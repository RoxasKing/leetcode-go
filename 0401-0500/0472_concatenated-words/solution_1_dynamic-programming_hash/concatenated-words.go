package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Dynamic Programming
// Hash

func findAllConcatenatedWordsInADict(words []string) []string {
	wordSet := map[string]struct{}{}
	sizeSet := map[int]struct{}{}
	for _, word := range words {
		wordSet[word] = struct{}{}
		sizeSet[len(word)] = struct{}{}
	}
	sizes := make([]int, 0, len(sizeSet))
	for size := range sizeSet {
		sizes = append(sizes, size)
	}
	sort.Ints(sizes)
	var out []string
	for _, word := range words {
		if len(word) == 0 {
			continue
		}
		delete(wordSet, word)
		n := len(word)
		f := make([]bool, n+1)
		f[0] = true
		for i := 1; i <= n; i++ {
			for _, size := range sizes {
				if i-size < 0 {
					break
				}
				if !f[i-size] {
					continue
				}
				if _, ok := wordSet[word[i-size:i]]; ok {
					f[i] = true
					break
				}
			}
		}
		if f[n] {
			out = append(out, word)
		}
		wordSet[word] = struct{}{}
	}
	return out
}
