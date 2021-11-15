package main

import "sort"

// Bit Maniputlation
// Sorting

func maxProduct(words []string) int {
	sort.Slice(words, func(i, j int) bool { return len(words[i]) > len(words[j]) })

	n := len(words)
	arr := make([]int, n)
	for j, w := range words {
		for i := range w {
			arr[j] |= 1 << int(w[i]-'a')
		}
	}

	out := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && len(words[i])*len(words[j]) > out; j++ {
			if arr[i]^arr[j] == arr[i]+arr[j] {
				out = Max(out, len(words[i])*len(words[j]))
			}
		}
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
