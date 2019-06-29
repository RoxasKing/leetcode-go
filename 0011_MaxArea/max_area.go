package main

import "fmt"

func maxArea(height []int) int {
	var max, tmp int
	var low, high = 0, len(height)-1

	for low < high {
		if height[low] < height[high] {
			tmp = height[low] * (high - low)
			low++
		} else {
			tmp = height[high] * (high - low)
			high--
		}
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func main() {
	slice := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(slice))
}
