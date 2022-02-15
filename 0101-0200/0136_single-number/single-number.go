package main

// Difficulty:
// Easy

// Tags:
// Bit Manipulation

func singleNumber(nums []int) int {
	out := 0
	for _, num := range nums {
		out ^= num
	}
	return out
}
