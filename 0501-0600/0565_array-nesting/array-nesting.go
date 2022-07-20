package main

// Difficulty:
// Medium

// Tags:
// Graph

func arrayNesting(nums []int) int {
	o := 0
	vis := make([]bool, len(nums))
	for i := range nums {
		c := 0
		for ; !vis[i]; i = nums[i] {
			vis[i] = true
			c++
		}
		if o < c {
			o = c
		}
	}
	return o
}
