package My_LeetCode_In_Go

import ()

/*
  给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
  注意：答案中不可以包含重复的三元组。
*/

func threeSum(nums []int) [][]int {
	quickSort(nums)
	var out [][]int
	var i int
	for i <= len(nums)-3 {
		head, tail := i+1, len(nums)-1
		for head < tail {
			sum := nums[i] + nums[head] + nums[tail]
			if sum == 0 {
				out = append(out, []int{nums[i], nums[head], nums[tail]})
				head++
				for head <= tail && nums[head] == nums[head-1] {
					head++
				}
				tail--
				for tail >= head && nums[tail] == nums[tail+1] {
					tail--
				}
			} else if sum < 0 {
				head++
			} else {
				tail--
			}
		}
		i++
		for i <= len(nums)-3 && nums[i] == nums[i-1] {
			i++
		}
	}
	return out
}
