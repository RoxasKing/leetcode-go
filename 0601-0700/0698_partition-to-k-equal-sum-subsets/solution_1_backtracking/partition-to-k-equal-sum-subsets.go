package main

import "sort"

// Difficulty:
// Hard

// Tags:
// Backtracking

func canPartitionKSubsets(nums []int, k int) bool {
	sort.Ints(nums)
	n := len(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 || nums[n-1] > sum/k {
		return false
	}
	avg := sum / k
	bucket := make([]int, k)
	for i := range bucket {
		bucket[i] = avg
	}
	return backtrack(nums, avg, bucket, n-1)
}

func backtrack(nums []int, avg int, bucket []int, idx int) bool {
	if idx < 0 {
		return true
	}
	for i := range bucket {
		if bucket[i] == nums[idx] || bucket[i]-nums[idx] >= nums[0] {
			bucket[i] -= nums[idx]
			if backtrack(nums, avg, bucket, idx-1) {
				return true
			}
			bucket[i] += nums[idx]
		}
		if bucket[i] == avg {
			break
		}
	}
	return false
}
