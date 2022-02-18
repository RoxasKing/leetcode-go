package main

// Difficulty:
// Medium

// Tags:
// Monotonic Stack

func removeKdigits(num string, k int) string {
	n := len(num)
	if n <= k {
		return "0"
	}
	nums := make([]byte, 0, n-k)
	for i := 0; i < n; i++ {
		for k > 0 && len(nums) > 0 && nums[len(nums)-1] > num[i] {
			k--
			nums = nums[:len(nums)-1]
		}
		nums = append(nums, num[i])
	}
	for nums = nums[:len(nums)-k]; len(nums) > 1 && nums[0] == '0'; nums = nums[1:] {
	}
	return string(nums)
}
