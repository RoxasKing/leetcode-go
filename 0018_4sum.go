package leetcode

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
	for i := 0; i <= len(nums)-4; {
		if nums[i] > target/4 {
			break
		}
		for j := i + 1; j <= len(nums)-3; {
			head, tail := j+1, len(nums)-1
			for head < tail {
				sum := nums[i] + nums[j] + nums[head] + nums[tail]
				if sum == target {
					out = append(
						out,
						[]int{nums[i], nums[j], nums[head], nums[tail]},
					)
					for head < tail && nums[head] == nums[head+1] {
						head++
					}
					head++
					for head < tail && nums[tail] == nums[tail-1] {
						tail--
					}
					tail--
				} else if sum < target {
					for head < tail && nums[head] == nums[head+1] {
						head++
					}
					head++
				} else {
					for head < tail && nums[tail] == nums[tail-1] {
						tail--
					}
					tail--
				}
			}
			for j < len(nums)-3 && nums[j] == nums[j+1] {
				j++
			}
			j++
		}
		for i < len(nums)-4 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return out
}
