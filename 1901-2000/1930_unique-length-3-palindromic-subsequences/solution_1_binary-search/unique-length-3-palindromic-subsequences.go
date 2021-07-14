package main

import "sort"

// Tags:
// Binary Search

func countPalindromicSubsequence(s string) int {
	dict := [26][]int{}
	for i := range s {
		dict[s[i]-'a'] = append(dict[s[i]-'a'], i)
	}

	var out int
	for i := 0; i < 26; i++ {
		if len(dict[i]) < 2 {
			continue
		}
		l, r := dict[i][0], dict[i][len(dict[i])-1]
		for j := 0; j < 26; j++ {
			if len(dict[j]) < 1 {
				continue
			}
			k := sort.Search(len(dict[j]), func(k int) bool { return dict[j][k] > l })
			if k < len(dict[j]) && l < dict[j][k] && dict[j][k] < r {
				out++
			}
		}
	}
	return out
}
