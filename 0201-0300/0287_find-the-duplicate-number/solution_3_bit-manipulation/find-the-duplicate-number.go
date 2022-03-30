package main

// Difficulty:
// Medium

// Tags:
// Bit Manipulation

func findDuplicate(nums []int) int {
	out := 0
	maxBit := 31
	for (len(nums)-1)>>maxBit == 0 {
		maxBit--
	}
	for bit := 0; bit <= maxBit; bit++ {
		var x, y int
		for i, num := range nums {
			if num&(1<<bit) != 0 {
				x++
			}
			if i > 0 && i&(1<<bit) != 0 {
				y++
			}
		}
		if x > y {
			out |= 1 << bit
		}
	}
	return out
}
