package main

// Tags:
// Dynamic Programming

func maxProduct(nums []int) int {
	max, min, out := nums[0], nums[0], nums[0]
	for _, num := range nums[1:] {
		a, b := max*num, min*num
		max = Max(num, Max(a, b))
		min = Min(num, Min(a, b))
		out = Max(out, max)
	}
	return out
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
