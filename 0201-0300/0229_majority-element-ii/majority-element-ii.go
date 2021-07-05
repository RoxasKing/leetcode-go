package main

import "sort"

// Sort

func majorityElement(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	k := n/3 + 1
	out := []int{}
	for l, r := 0, 0; r <= n; r++ {
		if r < n && nums[r] == nums[l] {
			continue
		}
		if r-l >= k {
			out = append(out, nums[l])
		}
		l = r
	}
	return out
}
