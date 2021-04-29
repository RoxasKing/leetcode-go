package main

/*
  A peak element is an element that is strictly greater than its neighbors.

  Given an integer array nums, find a peak element, and return its index. If the array contains multiple peaks, return the index to any of the peaks.

  You may imagine that nums[-1] = nums[n] = -∞.

  Example 1:
    Input: nums = [1,2,3,1]
    Output: 2
    Explanation: 3 is a peak element and your function should return the index number 2.

  Example 2:
    Input: nums = [1,2,1,3,5,6,4]
    Output: 5
    Explanation: Your function can return either index number 1 where the peak element is 2, or index number 5 where the peak element is 6.

  Constraints:
    1. 1 <= nums.length <= 1000
    2. -2^31 <= nums[i] <= 2^31 - 1
    3. nums[i] != nums[i + 1] for all valid i.

  Follow up: Could you implement a solution with logarithmic complexity?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-peak-element
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r {
		m := (l + r) >> 1
		if nums[m] <= nums[m+1] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
