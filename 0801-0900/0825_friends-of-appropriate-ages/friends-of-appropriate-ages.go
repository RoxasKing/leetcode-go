package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Two Pointers

func numFriendRequests(ages []int) int {
	sort.Ints(ages)
	var out, l, r int
	for _, age := range ages {
		if age < 15 {
			continue
		}
		for ; ages[l]<<1 <= age+14; l++ {
		}
		for ; r+1 < len(ages) && ages[r+1] <= age; r++ {
		}
		out += r - l
	}
	return out
}
