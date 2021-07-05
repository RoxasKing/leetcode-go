package main

func pivotIndex(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}

	sumR := make([]int, n)
	sumR[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		sumR[i] = sumR[i+1] + nums[i]
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
		if sum == sumR[i] {
			return i
		}
	}

	return -1
}
