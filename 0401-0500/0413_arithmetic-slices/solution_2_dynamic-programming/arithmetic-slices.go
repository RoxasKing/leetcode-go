package main

/*
  An integer array is called arithmetic if it consists of at least three elements and if the difference between any two consecutive elements is the same.

    For example, [1,3,5,7,9], [7,7,7,7], and [3,-1,-5,-9] are arithmetic sequences.

  Given an integer array nums, return the number of arithmetic subarrays of nums.

  A subarray is a contiguous subsequence of the array.

  Example 1:
    Input: nums = [1,2,3,4]
    Output: 3
    Explanation: We have 3 arithmetic slices in nums: [1, 2, 3], [2, 3, 4] and [1,2,3,4] itself.

  Example 2:
    Input: nums = [1]
    Output: 0

  Constraints:
    1. 1 <= nums.length <= 5000
    2. -1000 <= nums[i] <= 1000

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/arithmetic-slices
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Dynamic Programming

func numberOfArithmeticSlices(nums []int) int {
	n := len(nums)
	out, cnt := 0, 0
	for i := 2; i < n; i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			cnt++
			out += cnt
		} else {
			cnt = 0
		}
	}
	return out
}
