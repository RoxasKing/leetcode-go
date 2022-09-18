package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Hash

func palindromePairs(words []string) [][]int {
	type pair struct {
		i int
		w string
	}
	sz := map[int]struct{}{}
	a := []pair{}
	for i, w := range words {
		sz[len(w)] = struct{}{}
		a = append(a, pair{i, w})
	}
	sort.Slice(a, func(i, j int) bool {
		return len(a[i].w) < len(a[j].w) || len(a[i].w) == len(a[j].w) && a[i].i < a[j].i
	})
	revStr := func(s string) string {
		a := []byte(s)
		for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
			a[l], a[r] = a[r], a[l]
		}
		return string(a)
	}
	isPalindrome := func(s string) bool {
		for l, r := 0, len(s)-1; l < r; l, r = l+1, r-1 {
			if s[l] != s[r] {
				return false
			}
		}
		return true
	}
	h := map[string]int{}
	o := [][]int{}
	for _, p := range a {
		i, w := p.i, p.w
		for k := 0; k < len(w); k++ {
			if _, ok := sz[k]; !ok {
				continue
			}
			if j, ok := h[revStr(w[:k])]; ok && isPalindrome(w[k:]) {
				o = append(o, []int{i, j})
			}
			if j, ok := h[revStr(w[len(w)-k:])]; ok && isPalindrome(w[:len(w)-k]) {
				o = append(o, []int{j, i})
			}
		}
		if j, ok := h[revStr(w)]; ok {
			o = append(o, []int{j, i}, []int{i, j})
		}
		h[w] = i
	}
	return o
}
