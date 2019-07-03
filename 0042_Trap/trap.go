package main

import "fmt"

func trap(height []int) int {
	var out int
	size := len(height)
	limit := 1
	start := 0
	max := 0
	for {
		for i := 0; i < size; i++ {
			if (i == 0 || i > 0 && height[i-1] < limit) && height[i] >= limit {
				start = i
			} else if i > 1 && height[i-1] < limit && height[i] >= limit {
				out += i - start - 1
				start = i
			}
			if height[i] > max {
				max = height[i]
			}
		}
		if limit < max {
			limit++
		} else {
			break
		}
	}
	return out
}

func main() {
	in := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	fmt.Println(trap(in))
}
