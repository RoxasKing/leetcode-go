package main

// Tags:
// Bit Manipulation
func hammingWeight(num uint32) int {
	out := 0
	for num > 0 {
		out += int(num & 1)
		num >>= 1
	}
	return out
}
