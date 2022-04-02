package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Hash
// Sorting

func canReorderDoubled(arr []int) bool {
	sort.Ints(arr)
	freq := map[int]int{}
	cnt := 0
	for _, x := range arr {
		if freq[x<<1] > 0 {
			freq[x<<1]--
			cnt++
			continue
		}
		if x&1 == 0 && freq[x>>1] > 0 {
			freq[x>>1]--
			cnt++
			continue
		}
		freq[x]++
	}
	return cnt == len(arr)>>1
}
