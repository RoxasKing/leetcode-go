package main

import "sort"

/*
  Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

  Notice that the solution set must not contain duplicate triplets.

  Example 1:
    Input: nums = [-1,0,1,2,-1,-4]
    Output: [[-1,-1,2],[-1,0,1]]

  Example 2:
    Input: nums = []
    Output: []

  Example 3:
    Input: nums = [0]
    Output: []

  Constraints:
    1. 0 <= nums.length <= 3000
    2. -10^5 <= nums[i] <= 10^5

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/3sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	out := [][]int{}
	n := len(nums)
	for i := 0; i < n-2; i++ {
		j, k := i+1, n-1
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
		for i+1 < n-2 && nums[i+1] == nums[i] {
			i++
		}
	}
	return out
}
