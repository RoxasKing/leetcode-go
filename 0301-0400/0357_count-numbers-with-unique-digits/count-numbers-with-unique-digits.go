package main

// Difficulty:
// Medium

// Tags:
// Math

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	o := 10
	for x, i := 9, 0; i < n-1; i++ {
		x *= 9 - i
		o += x
	}
	return o
}
