package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Quick Select

func findKthLargest(nums []int, k int) int {
	for l, r := 0, len(nums)-1; l < r; {
		pivotIdx := rand.Intn(r+1-l) + l
		nums[pivotIdx], nums[r] = nums[r], nums[pivotIdx]
		i, j := l-1, r
		for {
			for i++; i < r && nums[i] > nums[r]; i++ {
			}
			for j--; j > l && nums[j] < nums[r]; j-- {
			}
			if i >= j {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		nums[i], nums[r] = nums[r], nums[i]
		if i < k-1 {
			l = i + 1
		} else if i > k-1 {
			r = i - 1
		} else {
			break
		}
	}
	return nums[k-1]
}
