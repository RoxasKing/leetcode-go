package main

import (
	"sort"
)

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	dict := make([]int64, n)
	dict[0] = 1
	for i := int64(1); i < int64(n); i++ {
		dict[i] = (dict[i-1] * 2) % divisor
	}
	var count int64
	l, r := 0, n-1
	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			count = (count + dict[r-l]) % divisor
			l++
		}
	}
	return int(count)
}

const divisor = 1e9 + 7
