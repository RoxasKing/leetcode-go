package main

import "math/rand"

/*
  Given an array of integers nums, sort the array in ascending order.

  Example 1:
    Input: nums = [5,2,3,1]
    Output: [1,2,3,5]

  Example 2:
    Input: nums = [5,1,1,2,0,0]
    Output: [0,0,1,1,2,5]

  Constraints:
    1. 1 <= nums.length <= 5 * 10^4
    2. -5 * 104 <= nums[i] <= 5 * 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/sort-an-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Quick Sort
func sortArray(nums []int) []int {
	qSort(nums)
	return nums
}

func qSort(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}
	pivotIdx := rand.Intn(n)
	nums[pivotIdx], nums[n-1] = nums[n-1], nums[pivotIdx]
	idx := 0
	for i := 0; i < n-1; i++ {
		if nums[i] <= nums[n-1] {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
		}
	}
	nums[idx], nums[n-1] = nums[n-1], nums[idx]
	qSort(nums[:idx])
	qSort(nums[idx+1:])
}
