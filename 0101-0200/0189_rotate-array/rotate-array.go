package main

func rotate(nums []int, k int) {
	k %= len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(nums []int) {
	n := len(nums)
	for i := 0; i < n>>1; i++ {
		nums[i], nums[n-1-i] = nums[n-1-i], nums[i]
	}
}
