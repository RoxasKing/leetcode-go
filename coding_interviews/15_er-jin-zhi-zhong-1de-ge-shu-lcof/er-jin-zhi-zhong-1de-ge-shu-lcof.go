package main

// Tags:
// Bit Manipulation

func hammingWeight(num uint32) int {
	out := 0
	for ; num > 0; num &= num - 1 {
		out++
	}
	return out
}
