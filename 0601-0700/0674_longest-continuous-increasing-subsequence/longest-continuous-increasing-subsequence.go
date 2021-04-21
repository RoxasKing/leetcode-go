package main

/*
  Given an unsorted array of integers nums, return the length of the longest continuous increasing subsequence (i.e. subarray). The subsequence must be strictly increasing.

  A continuous increasing subsequence is defined by two indices l and r (l < r) such that it is [nums[l], nums[l + 1], ..., nums[r - 1], nums[r]] and for each l <= i < r, nums[i] < nums[i + 1].

  Example 1:
    Input: nums = [1,3,5,4,7]
    Output: 3
    Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
    Even though [1,3,5,7] is an increasing subsequence, it is not continuous as elements 5 and 7 are separated by element 4.

  Example 2:
    Input: nums = [2,2,2,2,2]
    Output: 1
    Explanation: The longest continuous increasing subsequence is [2] with length 1. Note that it must be strictly
    increasing.

  Constraints:
    0 <= nums.length <= 10^4
    -10^9 <= nums[i] <= 10^9

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findLengthOfLCIS(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	max := 1
	cur := 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			cur++
		} else {
			max = Max(max, cur)
			cur = 1
		}
	}
	max = Max(max, cur)
	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
