package main

// Difficulty:
// Easy

// Tags:
// Counting

func specialArray(nums []int) int {
	h := [1001]int{}
	for _, x := range nums {
		h[x]++
	}
	for c, i := 0, 0; i < 1000; c, i = c+h[i], i+1 {
		if len(nums)-c == i {
			return i
		}
	}
	return -1
}
