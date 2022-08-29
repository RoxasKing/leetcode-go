package main

// Difficulty:
// Easy

func maxProduct(nums []int) int {
	a, b := -1, -1
	for _, x := range nums {
		if a < x {
			a, b = x, a
		} else if b < x {
			b = x
		}
	}
	return (a - 1) * (b - 1)
}
