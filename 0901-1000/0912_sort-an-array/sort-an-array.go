package main

import "math/rand"

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
