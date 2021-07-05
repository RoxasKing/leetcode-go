package main

import (
	"strconv"
)

// Tags:
// Bit Manipulation
// Math

func findNthDigit(n int) int {
	b, i := 1, 1
	for ; n > i*9*b; b, i = b*10, i+1 {
		n -= i * 9 * b
	}
	x := n / i
	n -= x * i
	if n == 0 {
		return int(strconv.Itoa(b + x - 1)[i-1] - '0')
	}
	return int(strconv.Itoa(b + x)[n-1] - '0')
}
