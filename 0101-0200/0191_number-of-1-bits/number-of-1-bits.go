package main

// Difficulty:
// Easy

// Tags:
// Bit Manipulation

func hammingWeight(num uint32) int {
	o := 0
	for ; num > 0; num &= num - 1 {
		o++
	}
	return o
}
