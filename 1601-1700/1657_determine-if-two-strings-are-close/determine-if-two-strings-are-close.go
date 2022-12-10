package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Sorting

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	h1, h2 := [26]int{}, [26]int{}
	e1, e2 := [26]bool{}, [26]bool{}
	for _, c := range word1 {
		h1[c-'a']++
		e1[c-'a'] = true
	}
	for _, c := range word2 {
		h2[c-'a']++
		e2[c-'a'] = true
	}
	a1 := make([]int, 26)
	a2 := make([]int, 26)
	for i := range h1 {
		a1[i] = h1[i]
		a2[i] = h2[i]
	}
	sort.Ints(a1)
	sort.Ints(a2)
	for i := 0; i < 26; i++ {
		if e1[i] && !e2[i] || !e1[i] && e2[i] || a1[i] != a2[i] {
			return false
		}
	}
	return true
}
