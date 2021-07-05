package main

import "sort"

// Binary Search
func numSmallerByFrequency(queries []string, words []string) []int {
	m, n := len(queries), len(words)
	arr := make([]int, n)
	for i, word := range words {
		cnt := [26]int{}
		for j := range word {
			cnt[word[j]-'a']++
		}
		for j := 0; j < 26; j++ {
			if cnt[j] > 0 {
				arr[i] = cnt[j]
				break
			}
		}
	}
	sort.Ints(arr)

	qs := make([]int, m)
	for i, query := range queries {
		cnt := [26]int{}
		for j := range query {
			cnt[query[j]-'a']++
		}
		for j := 0; j < 26; j++ {
			if cnt[j] > 0 {
				qs[i] = cnt[j]
				break
			}
		}
	}

	out := make([]int, m)
	for i, q := range qs {
		j := sort.Search(n, func(j int) bool { return q < arr[j] })
		out[i] = n - j
	}
	return out
}
