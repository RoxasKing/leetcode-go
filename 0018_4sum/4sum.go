package main

import (
	"sort"
)

/*
  给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
  注意：
  答案中不可以包含重复的四元组。
*/

func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	var out [][]int
	for a := 0; a < len(nums)-3; a++ {
		if nums[a] > target/4 {
			break
		}
		for b := a + 1; b < len(nums)-2; b++ {
			if nums[a]+nums[b] > target/2 {
				break
			}
			c, d := b+1, len(nums)-1
			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum < target {
					for c+1 < d && nums[c+1] == nums[c] {
						c++
					}
					c++
				} else if sum > target {
					for c < d-1 && nums[d] == nums[d-1] {
						d--
					}
					d--
				} else {
					out = append(out, []int{nums[a], nums[b], nums[c], nums[d]})
					for c+1 < d && nums[c+1] == nums[c] {
						c++
					}
					c++
					for c < d-1 && nums[d] == nums[d-1] {
						d--
					}
					d--
				}
			}
			for b+1 < len(nums)-2 && nums[b+1] == nums[b] {
				b++
			}
		}
		for a+1 < len(nums)-3 && nums[a+1] == nums[a] {
			a++
		}
	}
	return out
}
