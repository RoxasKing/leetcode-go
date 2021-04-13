package main

/*
  There is an integer array nums sorted in non-decreasing order (not necessarily with distinct values).

  Before being passed to your function, nums is rotated at an unknown pivot index k (0 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,4,4,5,6,6,7] might be rotated at pivot index 5 and become [4,5,6,6,7,0,1,2,4,4].

  Given the array nums after the rotation and an integer target, return true if target is in nums, or false if it is not in nums.

  Example 1:
    Input: nums = [2,5,6,0,0,1,2], target = 0
    Output: true

  Example 2:
    Input: nums = [2,5,6,0,0,1,2], target = 3
    Output: false

  Constraints:
    1. 1 <= nums.length <= 5000
    2. -10^4 <= nums[i] <= 10^4
    3. nums is guaranteed to be rotated at some pivot.
    4. -10^4 <= target <= 10^4

  Follow up: This problem is the same as Search in Rotated Sorted Array, where nums may contain duplicates. Would this affect the runtime complexity? How and why?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func search(nums []int, target int) bool {
	l, r := 0, len(nums)-1
	for l < r && nums[l] == nums[r] {
		l++
	}
	rotateIdx := bSearchRotateIdx(nums, l, r)
	if rotateIdx == l || nums[l] > target {
		return bSearch(nums, rotateIdx, r, target)
	}
	return bSearch(nums, l, rotateIdx-1, target)
}

func bSearchRotateIdx(nums []int, l, r int) int {
	start := l
	for l < r {
		m := (l + r) >> 1
		if nums[m] > nums[m+1] {
			return m + 1
		}
		if nums[l] <= nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return start
}

func bSearch(nums []int, l, r, target int) bool {
	for l <= r {
		m := (l + r) >> 1
		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return true
		}
	}
	return false
}
