package main

import "sort"

/*
  给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
*/
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	out := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > target/3 {
			break
		}
		l, r := i+1, len(nums)-1
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
