package My_LeetCode_In_Go

import "math/rand"

func findKthLargest_heap(nums []int, k int) int {
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

func findKthLargest(nums []int, k int) int {
	if len(nums) < 2 {
		return nums[0]
	}
	pivotIndex := rand.Intn(len(nums))
	pivot := nums[pivotIndex]
	nums[0], nums[pivotIndex] = nums[pivotIndex], nums[0]
	mid, tail := 0, len(nums)-1
	for i := 1; i <= tail; {
		if nums[i] > pivot {
			nums[i], nums[tail], tail = nums[tail], nums[i], tail-1
		} else {
			nums[i], nums[mid], mid, i = nums[mid], nums[i], mid+1, i+1
		}
	}
	if mid < len(nums)-k {
		return findKthLargest(nums[mid+1:], k)
	} else if mid > len(nums)-k {
		return findKthLargest(nums[:mid], k-(len(nums)-mid))
	}
	return pivot
}
