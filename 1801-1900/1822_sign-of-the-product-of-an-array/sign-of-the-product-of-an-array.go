package main

// Difficulty:
// Easy

func arraySign(nums []int) int {
	o := 1
	for _, x := range nums {
		if x == 0 {
			return 0
		} else if x < 0 {
			o = -o
		}
	}
	return o
}
