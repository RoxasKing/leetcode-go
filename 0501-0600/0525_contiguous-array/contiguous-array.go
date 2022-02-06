package main

// Difficulty:
// Medium

// Tags:
// Hash

func findMaxLength(nums []int) int {
	out := 0
	h := map[int]int{0: -1}
	c := 0
	for i, num := range nums {
		if num == 0 {
			c--
		} else {
			c++
		}
		if j, ok := h[c]; !ok {
			h[c] = i
		} else if out < i-j {
			out = i - j
		}
	}
	return out
}
