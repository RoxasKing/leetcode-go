package main

import "sort"

// Tags:
// Two Pointers
// Sorting

func countPairs(deliciousness []int) int {
	cnt := map[int]int{}
	for _, d := range deliciousness {
		cnt[d]++
	}
	arr := make([]int, 0, len(cnt))
	for x := range cnt {
		arr = append(arr, x)
	}
	sort.Ints(arr)

	out := 0
	for good := 1; good <= 1<<21; good <<= 1 {
		for l, r := 0, len(arr)-1; l <= r; {
			sum := arr[l] + arr[r]
			if sum < good {
				l++
				continue
			} else if sum > good {
				r--
				continue
			}
			if l == r {
				out += cnt[arr[l]] * (cnt[arr[l]] - 1) >> 1
			} else {
				out += cnt[arr[l]] * cnt[arr[r]]
			}
			out %= 1e9 + 7
			l++
			r--
		}
	}
	return out
}
