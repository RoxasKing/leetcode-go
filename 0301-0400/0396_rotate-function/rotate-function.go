package main

// Tags:
// Math

func maxRotateFunction(nums []int) int {
	n := len(nums)
	sum := 0
	f := 0
	for i, num := range nums {
		sum += num
		f += i * num
	}

	out := f
	for i := 1; i < n; i++ {
		f -= sum - nums[i-1]
		f += nums[i-1] * (n - 1)
		out = Max(out, f)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
