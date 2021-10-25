package main

import "math"

// Difficulty:
// Easy

// Tags:
// Math

func constructRectangle(area int) []int {
	for i := int(math.Sqrt(float64(area))); i > 0; i-- {
		if area%i == 0 {
			return []int{area / i, i}
		}
	}
	return []int{area, 1}
}
