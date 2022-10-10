package main

import "sort"

// Difficulty:
// Medium

// Tags:
// Two Pointers

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	o := 1<<31 - 1
	for i := 0; i < n-2; i++ {
		for l, r := i+1, n-1; l < r; {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return sum
			} else if sum < target {
				l++
			} else {
				r--
			}
			if abs(o-target) > abs(sum-target) {
				o = sum
			}
		}
	}
	return o
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
