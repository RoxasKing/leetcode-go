package main

// Tags:
// Recursion
func PredictTheWinner(nums []int) bool {
	return maxScore(nums, 1) >= 0
}

func maxScore(nums []int, turn int) int {
	if len(nums) == 1 {
		return nums[0] * turn
	}
	head, tail := 0, len(nums)-1
	pickHead := nums[head]*turn + maxScore(nums[head+1:], -turn)
	pickTail := nums[tail]*turn + maxScore(nums[:tail], -turn)
	return Max(pickHead*turn, pickTail*turn) * turn
	// when turn == -1, get min score
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Dynamic Programming
func PredictTheWinner2(nums []int) bool {
	n := len(nums)
	dp := make([]int, n)
	copy(dp, nums)
	for i := n - 2; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[j] = Max(nums[i]-dp[j], nums[j]-dp[j-1])
		}
	}
	return dp[n-1] >= 0
}
