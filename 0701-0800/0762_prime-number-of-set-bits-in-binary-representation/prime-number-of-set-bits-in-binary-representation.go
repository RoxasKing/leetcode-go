package main

import "math/bits"

// Difficulty:
// Easy

// Tags:
// Math

func countPrimeSetBits(left int, right int) int {
	out := 0
	for x := left; x <= right; x++ {
		if 1<<bits.OnesCount(uint(x))&0b10100010100010101100 != 0 {
			out++
		}
	}
	return out
}
