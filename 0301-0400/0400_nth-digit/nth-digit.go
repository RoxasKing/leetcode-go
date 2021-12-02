package main

import "strconv"

// Difficulty:
// Medium

// Tags:
// Math

func findNthDigit(n int) int {
	size, base := 1, 1
	for n > size*9*base {
		n -= size * 9 * base
		size++
		base *= 10
	}
	x := n / size
	n -= x * size
	if n == 0 {
		x--
		n = size
	}
	return int(strconv.Itoa(base + x)[n-1] - '0')
}

// 1*9 1~9
// 2*90 10~99
// 3*900 100~999
// 4*9000 1000~9999
