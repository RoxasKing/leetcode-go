package main

/*
  Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:
    1. [4,5,6,7,0,1,2] if it was rotated 4 times.
    2. [0,1,2,4,5,6,7] if it was rotated 7 times.
  Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

  Given the sorted rotated array nums of unique elements, return the minimum element of this array.

  Example 1:
    Input: nums = [3,4,5,1,2]
    Output: 1
    Explanation: The original array was [1,2,3,4,5] rotated 3 times.

  Example 2:
    Input: nums = [4,5,6,7,0,1,2]
    Output: 0
    Explanation: The original array was [0,1,2,4,5,6,7] and it was rotated 4 times.

  Example 3:
    Input: nums = [11,13,15,17]
    Output: 11
    Explanation: The original array was [11,13,15,17] and it was rotated 4 times.

  Constraints:
    1. n == nums.length
    2. 1 <= n <= 5000
    3. -5000 <= nums[i] <= 5000
    4. All the integers of nums are unique.
    5. nums is sorted and rotated between 1 and n times.

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Binary Search
func findMin(nums []int) int {
	return nums[bSearchRotateIdx(nums, 0, len(nums)-1)]
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
