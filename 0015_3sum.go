package My_LeetCode_In_Go

import "sort"

/*
  给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
  注意：答案中不可以包含重复的三元组。
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var out [][]int
	var i int
	for i <= len(nums)-3 {
		if nums[i] > 0 {
			break
		}
		head, tail := i+1, len(nums)-1
		for head < tail {
			sum := nums[i] + nums[head] + nums[tail]
			if sum == 0 {
				out = append(out, []int{nums[i], nums[head], nums[tail]})
				for head < tail && nums[head] == nums[head+1] {
					head++
				}
				for head < tail && nums[tail] == nums[tail-1] {
					tail--
				}
				head++
				tail--
			} else if sum < 0 {
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
		for i < len(nums)-3 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return out
}
