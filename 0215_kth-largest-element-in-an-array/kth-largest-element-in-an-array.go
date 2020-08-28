package main

import "math/rand"

/*
  在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

  说明:
    你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
*/

// Quick Sort
func findKthLargest(nums []int, k int) int {
	pivotIndex, lastIndex := rand.Intn(len(nums)), len(nums)-1
	nums[pivotIndex], nums[lastIndex] = nums[lastIndex], nums[pivotIndex]
	var index int
	for i := 0; i < len(nums); i++ {
		if nums[i] > nums[lastIndex] {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[index], nums[lastIndex] = nums[lastIndex], nums[index]
	if index+1 < k {
		return findKthLargest(nums[index+1:], k-index-1)
	} else if index+1 > k {
		return findKthLargest(nums[:index], k)
	}
	return nums[index]
}
