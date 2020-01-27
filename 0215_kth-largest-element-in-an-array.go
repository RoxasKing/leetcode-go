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
	nums[len(nums)-1], nums[pivotIndex] = nums[pivotIndex], nums[len(nums)-1]
	var mid int
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] < pivot {
			nums[i], nums[mid] = nums[mid], nums[i]
			mid++
		}
	}
	nums[mid], nums[len(nums)-1] = nums[len(nums)-1], nums[mid]
	if k < len(nums)-mid {
		return findKthLargest(nums[mid+1:], k)
	} else if k > len(nums)-mid {
		return findKthLargest(nums[:mid], k-(len(nums)-mid))
	}
	return pivot
}
