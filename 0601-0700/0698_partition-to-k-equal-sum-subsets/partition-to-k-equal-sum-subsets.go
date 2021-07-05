package main

import (
	"sort"
)

// Tags:
// Backtracking
func canPartitionKSubsets(nums []int, k int) bool {
	n := len(nums)
	sort.Ints(nums)
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k > 0 || nums[n-1] > sum/k {
		return false
	}

	target := sum / k
	bucket := make([]int, k)
	for i := range bucket {
		bucket[i] = target
	}

	return dfs(nums, target, bucket, n-1)
}

func dfs(nums []int, target int, bucket []int, idx int) bool {
	if idx < 0 {
		return true
	}

	for i := range bucket {
		if bucket[i] == nums[idx] || bucket[i]-nums[idx] >= nums[0] {
			bucket[i] -= nums[idx]
			if dfs(nums, target, bucket, idx-1) {
				return true
			}
			bucket[i] += nums[idx]
		}
		if bucket[i] == target {
			break
		}
	}

	return false
}

// Dynamic Programming
func canPartitionKSubsets2(nums []int, k int) bool {
	n := len(nums)

	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}

	target := sum / k

	f := make(map[int]bool)
	f[0] = true
	for i := 1; i < 1<<n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if i>>j&1 == 1 {
				sum += nums[j]
			}
		}
		if target > 0 && sum%target > 0 {
			continue
		}
		if sum == target {
			f[i] = true
		} else {
			for k := i; k > 0; k = (k - 1) & i {
				if !f[k] {
					continue
				}
				if f[k] && f[i^k] {
					f[i] = true
					break
				}
			}
		}
		if f[i] && f[(1<<n-1)^i] {
			return true
		}
	}
	return false
}
