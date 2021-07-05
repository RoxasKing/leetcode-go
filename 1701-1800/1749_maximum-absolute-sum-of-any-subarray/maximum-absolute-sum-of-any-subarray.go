package main

// Tags:
// Greedy Algoritm

func maxAbsoluteSum(nums []int) int {
	out := 0
	max, min := -1<<31, 1<<31-1
	for _, num := range nums {
		max = Max(max+num, num)
		min = Min(min+num, num)
		out = Max(out, Max(Abs(max), Abs(min)))
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

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
