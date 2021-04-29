package main

import "math/rand"

/*
  Given an integer array nums and an integer k, return the kth largest element in the array.

  Note that it is the kth largest element in the sorted order, not the kth distinct element.

  Example 1:
    Input: nums = [3,2,1,5,6,4], k = 2
    Output: 5

  Example 2:
    Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
    Output: 4

  Constraints:
    1. 1 <= k <= nums.length <= 10^4
    2. -10^4 <= nums[i] <= 10^4

  来源：力扣（LeetCode）
  链接：https://leetcode-cn.com/problems/kth-largest-element-in-an-array
  著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// Quick Sort
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	pivotIdx := rand.Intn(n)
	nums[pivotIdx], nums[n-1] = nums[n-1], nums[pivotIdx]
	i := 0
	for j := 0; j < n-1; j++ {
		if nums[j] > nums[n-1] {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[n-1] = nums[n-1], nums[i]
	if i+1 < k {
		return findKthLargest(nums[i+1:], k-(i+1))
	} else if i+1 > k {
		return findKthLargest(nums[:i], k)
	}
	return nums[i]
}
