package leetcode

import "math/rand"

/*
  在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

  说明:
    你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
*/

// Quick Sort
func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	pivotIndex := rand.Intn(len(nums))
	pivot := nums[pivotIndex]
	nums[len(nums)-1], nums[pivotIndex] = nums[pivotIndex], nums[len(nums)-1]
	var index int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < pivot {
			nums[i], nums[index] = nums[index], nums[i]
			index++
		}
	}
	nums[index], nums[len(nums)-1] = nums[len(nums)-1], nums[index]
	if k < len(nums)-index {
		return findKthLargest(nums[index+1:], k)
	} else if k > len(nums)-index {
		return findKthLargest(nums[:index], k-(len(nums)-index))
	}
	return pivot
}

// Heap Sort
func findKthLargest2(nums []int, k int) int {
	nodeRise := func() {
		for i := k/2 - 1; i >= 0; i-- {
			son := i*2 + 1
			if son > k-1 {
				return
			}
			if son < k-1 && nums[son+1] < nums[son] {
				son++
			}
			if nums[son] < nums[i] {
				nums[son], nums[i] = nums[i], nums[son]
			}
		}
	}
	nodeRise()
	for i := k; i < len(nums); i++ {
		if nums[i] > nums[0] {
			nums[0], nums[i] = nums[i], nums[0]
			nodeRise()
		}
	}
	return nums[0]
}
