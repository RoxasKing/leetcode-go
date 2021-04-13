package main

import "sort"

/*
  给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

  注意：答案中不可以包含重复的三元组。

  示例：
    给定数组 nums = [-1, 0, 1, 2, -1, -4]，
    满足要求的三元组集合为：
    [
      [-1, 0, 1],
      [-1, -1, 2]
    ]

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/3sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var out [][]int
	for i := 0; i < len(nums)-2; i++ {
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				for j+1 < k && nums[j+1] == nums[j] {
					j++
				}
				j++
			} else if sum > 0 {
				for j < k-1 && nums[k-1] == nums[k] {
					k--
				}
				k--
			} else {
				out = append(out, []int{nums[i], nums[j], nums[k]})
				for j+1 < k && nums[j+1] == nums[j] {
					j++
				}
				j++
				for j < k-1 && nums[k-1] == nums[k] {
					k--
				}
				k--
			}
		}
		for i+1 < len(nums)-2 && nums[i+1] == nums[i] {
			i++
		}
	}
	return out
}
