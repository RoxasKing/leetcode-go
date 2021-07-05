package main

import "sort"

// Two Pointers
func countPairs(deliciousness []int) int {
	arr := []int{}
	cnt := make(map[int]int)
	for _, d := range deliciousness {
		if cnt[d] == 0 {
			arr = append(arr, d)
		}
		cnt[d]++
	}
	sort.Ints(arr)

	out := 0
	for good := 1; good <= 1<<21; good <<= 1 {
		for l, r := 0, len(arr)-1; l <= r; {
			sum := arr[l] + arr[r]
			if sum < good {
				l++
			} else if sum > good {
				r--
			} else {
				if l == r {
					out = (out + cnt[arr[l]]*(cnt[arr[l]]-1)/2) % (1e9 + 7)
				} else {
					out = (out + cnt[arr[l]]*cnt[arr[r]]) % (1e9 + 7)
				}
				l++
				r--
			}
		}
	}
	return out
}
