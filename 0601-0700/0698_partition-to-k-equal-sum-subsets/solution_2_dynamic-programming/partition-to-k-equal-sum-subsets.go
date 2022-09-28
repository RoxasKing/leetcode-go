package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming
// Bitmask

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	avg := sum / k
	f := map[int]bool{}
	n := len(nums)
	for i := 1; i < 1<<n; i++ {
		sum := 0
		for j := 0; j < n; j++ {
			if i>>(n-1-j)&1 == 1 {
				sum += nums[j]
			}
		}
		if sum%avg > 0 {
			continue
		}
		if sum == avg {
			f[i] = true
		} else {
			for j := i; j > 0; j = (j - 1) & i {
				if f[j] && f[i^j] {
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
