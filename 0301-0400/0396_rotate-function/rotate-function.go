package main

// Difficulty:
// Medium

// Tags:
// Math

func maxRotateFunction(nums []int) int {
	val, sum := 0, 0
	for i, x := range nums {
		val += i * x
		sum += x
	}
	n := len(nums)
	o := val
	for i := 1; i < n; i++ {
		if val += nums[i-1]*n - sum; o < val {
			o = val
		}
	}
	return o
}
