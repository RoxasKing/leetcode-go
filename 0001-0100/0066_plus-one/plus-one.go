package main

// Difficulty:
// Easy

func plusOne(digits []int) []int {
	n := len(digits)
	remain := 1
	for i := n - 1; i >= 0 && remain > 0; i-- {
		digits[i] += remain
		remain = digits[i] / 10
		digits[i] %= 10
	}
	if remain > 0 {
		tmp := make([]int, n+1)
		copy(tmp[1:], digits)
		tmp[0] = remain
		digits = tmp
	}
	return digits
}
