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
	a := make([]int, n)
	for i := range a {
		x := 0
		for _, c := range words[i] {
			x |= 1 << int(c-'a')
		}
		a[i] = x
	}
	o := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n && o < len(words[i])*len(words[j]); j++ {
			if a[i]&a[j] == 0 {
				o = len(words[i]) * len(words[j])
			}
		}
	}
	return o
}
