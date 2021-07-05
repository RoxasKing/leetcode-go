package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return nil
	}
	rev := make([]int, n)
	rev[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		rev[i] = rev[i+1] * nums[i+1]
	}
	pos := 1
	out := make([]int, n)
	for i := range out {
		out[i] = pos * rev[i]
		pos *= nums[i]
	}
	return out
}
