package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Sorting
// Hash

func longestStrChain(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) < len(words[j]) })
	cnt := map[string]int{}
	o := 1
	for _, w := range words {
		cnt[w] = 1
		for i := 0; i < len(w); i++ {
			if v, ok := cnt[w[:i]+w[i+1:]]; ok {
				cnt[w] = max(cnt[w], v+1)
				o = max(o, cnt[w])
			}
		}
	}
	return o
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
