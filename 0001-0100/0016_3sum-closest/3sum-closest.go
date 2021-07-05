package main

import "sort"

func threeSumClosest(nums []int, target int) int {
	n := len(nums)
	sort.Ints(nums)
	out := nums[0] + nums[1] + nums[2]
	for i := 0; i < n-2 && nums[i] <= target/3; i++ {
		l, r := i+1, n-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum < target {
				for l+1 < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			} else if sum > target {
				for l < r-1 && nums[r-1] == nums[r] {
					r--
				}
				r--
			} else {
				return sum
			}
			if Abs(sum-target) < Abs(out-target) {
				out = sum
			}
		}
		for i+1 < len(nums)-2 && nums[i] == nums[i+1] {
			i++
		}
	}
	return out
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
