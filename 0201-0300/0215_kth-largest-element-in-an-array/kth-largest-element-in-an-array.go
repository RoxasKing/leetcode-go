package main

import "math/rand"

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
