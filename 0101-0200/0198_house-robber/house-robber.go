package main

// Difficulty:
// Medium

// Tags:
// Dynamic Programming

func rob(nums []int) int {
	f0, f1 := 0, 0
	for _, num := range nums {
		t := f0
		f0 = f1
		if t+num > f1 {
			f1 = t + num
		}
	}
	return f1
}
