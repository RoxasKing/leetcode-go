package main

// Tags:
// Binary Search

func smallestDivisor(nums []int, threshold int) int {
	l, r := 1, int(1e6)
	for l < r {
		m := (l + r) >> 1
		if getSum(nums, m) > threshold {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func getSum(nums []int, divisor int) int {
	out := 0
	for _, num := range nums {
		out += num / divisor
		if num%divisor > 0 {
			out++
		}
	}
	return out
}
