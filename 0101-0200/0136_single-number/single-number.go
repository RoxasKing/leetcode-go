package main

// Tags:
// Bit Manipulation
func singleNumber(nums []int) int {
	out := 0
	for _, num := range nums {
		out ^= num
	}
	return out
}
