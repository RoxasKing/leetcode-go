package main

import (
	"sort"
)

/*
  Given an array nums of n integers, return an array of all the unique quadruplets [nums[a], nums[b], nums[c], nums[d]] such that:
    1. 0 <= a, b, c, d < n
    2. a, b, c, and d are distinct.
    3. nums[a] + nums[b] + nums[c] + nums[d] == target
  You may return the answer in any order.

  Example 1:
    Input: nums = [1,0,-1,0,-2,2], target = 0
    Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

  Example 2:
    Input: nums = [2,2,2,2,2], target = 8
    Output: [[2,2,2,2]]

  Constraints:
    1. 1 <= nums.length <= 200
    2. -10^9 <= nums[i] <= 10^9
    3. -10^9 <= target <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/4sum
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Two Pointers
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	out := [][]int{}
	for a := 0; a < n-3 && nums[a] <= target/4; a++ {
		for b := a + 1; b < n-2 && nums[a]+nums[b] <= target/2; b++ {
			c, d := b+1, n-1
			for c < d {
				sum := nums[a] + nums[b] + nums[c] + nums[d]
				if sum < target {
					for c+1 < d && nums[c+1] == nums[c] {
						c++
					}
					c++
				} else if sum > target {
					for c < d-1 && nums[d-1] == nums[d] {
						d--
					}
					d--
				} else {
					out = append(out, []int{nums[a], nums[b], nums[c], nums[d]})
					for c+1 < d && nums[c+1] == nums[c] {
						c++
					}
					c++
					for c < d-1 && nums[d-1] == nums[d] {
						d--
					}
					d--
				}
			}
			for b+1 < n-2 && nums[b+1] == nums[b] {
				b++
			}
		}
		for a+1 < n-3 && nums[a+1] == nums[a] {
			a++
		}
	}
	return out
}
