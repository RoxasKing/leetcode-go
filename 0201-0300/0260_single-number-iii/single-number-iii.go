package main

import "math/bits"

// Difficulty:
// Medium

// Tags:
// Bit Manipulation

func singleNumber(nums []int) []int {
	xor := 0
	for _, num := range nums {
		xor ^= num
	}
	mask := 1 << bits.TrailingZeros(uint(xor))
	num1, num2 := 0, 0
	for _, num := range nums {
		if num&mask > 0 {
			num1 ^= num
		} else {
			num2 ^= num
		}
	}
	return []int{num1, num2}
}
