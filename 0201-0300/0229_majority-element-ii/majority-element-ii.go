package main

import "sort"

/*
  Given an integer array of size n, find all elements that appear more than ⌊ n/3 ⌋ times.

  Follow-up: Could you solve the problem in linear time and in O(1) space?

  Example 1:
    Input: nums = [3,2,3]
    Output: [3]

  Example 2:
    Input: nums = [1]
    Output: [1]

  Example 3:
    Input: nums = [1,2]
    Output: [1,2]

  Constraints:
    1. 1 <= nums.length <= 5 * 10^4
    2. -10^9 <= nums[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/majority-element-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Sort

func majorityElement(nums []int) []int {
	sort.Ints(nums)
	n := len(nums)
	k := n/3 + 1
	out := []int{}
	for l, r := 0, 0; r <= n; r++ {
		if r < n && nums[r] == nums[l] {
			continue
		}
		if r-l >= k {
			out = append(out, nums[l])
		}
		l = r
	}
	return out
}
