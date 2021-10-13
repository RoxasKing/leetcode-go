package main

// Difficulty:
// Medium

// Tags:
// Bit Manipulation

func rangeBitwiseAnd(left int, right int) int {
	k := 0
	for left < right {
		left >>= 1
		right >>= 1
		k++
	}
	return left << k
}
