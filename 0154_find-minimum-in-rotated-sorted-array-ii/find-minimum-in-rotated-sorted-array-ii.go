package main

/*
  Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,4,4,5,6,7] might become:
    1. [4,5,6,7,0,1,4] if it was rotated 4 times.
    2. [0,1,4,4,5,6,7] if it was rotated 7 times.
  Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

  Given the sorted rotated array nums that may contain duplicates, return the minimum element of this array.

  Example 1:
    Input: nums = [1,3,5]
    Output: 1

  Example 2:
    Input: nums = [2,2,2,0,1]
    Output: 0

  Constraints:
    1. n == nums.length
    2. 1 <= n <= 5000
    3. -5000 <= nums[i] <= 5000
    4. nums is sorted and rotated between 1 and n times.

  Follow up: This is the same as Find Minimum in Rotated Sorted Array but with duplicates. Would allow duplicates affect the run-time complexity? How and why?

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	for l < r && nums[l] == nums[r] {
		l++
	}
	return bSearchRotateNum(nums, l, r)
}

func bSearchRotateNum(nums []int, l, r int) int {
	start := l
	for l < r {
		m := (l + r) >> 1
		if nums[m] > nums[m+1] {
			return nums[m+1]
		}
		if nums[l] <= nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return nums[start]
}
