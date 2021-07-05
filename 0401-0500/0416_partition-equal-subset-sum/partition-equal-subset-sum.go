package main

// Tags:
// Dynamic Programming

func canPartition(nums []int) bool {
	n := len(nums)
	if n < 2 {
		return false
	}

	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum&1 == 1 {
		return false
	}

	m := sum >> 1
	dp := make([]bool, m+1)
	dp[0] = true

	for i := 0; i < n; i++ {
		for j := m; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}

	return dp[m]
}
