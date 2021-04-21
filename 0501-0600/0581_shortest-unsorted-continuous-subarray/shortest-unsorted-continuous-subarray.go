package main

import "sort"

/*
  Given an integer array nums, you need to find one continuous subarray that if you only sort this subarray in ascending order, then the whole array will be sorted in ascending order.

  Return the shortest such subarray and output its length.

  Example 1:
    Input: nums = [2,6,4,8,10,9,15]
    Output: 5
    Explanation: You need to sort [6, 4, 8, 10, 9] in ascending order to make the whole array sorted in ascending order.

  Example 2:
    Input: nums = [1,2,3,4]
    Output: 0

  Example 3:
    Input: nums = [1]
    Output: 0

  Constraints:
    1. 1 <= nums.length <= 10^4
    2. -10^5 <= nums[i] <= 10^5

  Follow up: Can you solve it in O(n) time complexity?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Ints(sorted)
	l, r := 0, n-1
	for l < n && sorted[l] == nums[l] {
		l++
	}
	for r >= l && sorted[r] == nums[r] {
		r--
	}
	return r + 1 - l
}

// Two Pointers
func findUnsortedSubarray2(nums []int) int {
	n := len(nums)
	min, max := 1<<31-1, -1<<31
	for i := 0; i < n-1; i++ {
		if nums[i] > nums[i+1] {
			max = Max(max, nums[i])
			min = Min(min, nums[i+1])
		}
	}

	l, r := 0, 0
	for i := 0; i < n; i++ {
		if nums[i] > min {
			l = i
			break
		}
	}
	for i := n - 1; i >= 0; i-- {
		if nums[i] < max {
			r = i
			break
		}
	}

	if l >= r {
		return 0
	}
	return r + 1 - l
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
