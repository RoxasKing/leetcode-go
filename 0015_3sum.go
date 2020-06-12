package leetcode

import "sort"

/*
  给定一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？找出所有满足条件且不重复的三元组。
  注意：答案中不可以包含重复的三元组。
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var out [][]int
	var l, r int
	for i := 0; i < len(nums)-2 && nums[i] <= 0; i++ {
		l, r = i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			switch {
			case sum < 0:
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
			case sum > 0:
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			default:
				out = append(out, []int{nums[i], nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				l++
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				r--
			}
		}
		for i < len(nums)-3 && nums[i] == nums[i+1] {
			i++
		}
	}
	return out
}
