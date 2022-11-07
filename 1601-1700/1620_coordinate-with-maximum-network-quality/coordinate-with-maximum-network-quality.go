package main

import "math"

// Difficulty:
// Medium

func bestCoordinate(towers [][]int, radius int) []int {
	o, x := []int{-1, -1}, -1
	for i := 0; i <= 50; i++ {
		for j := 0; j <= 50; j++ {
			sum := 0
			for _, e := range towers {
				if d2 := (i-e[0])*(i-e[0]) + (j-e[1])*(j-e[1]); d2 <= radius*radius {
					sum += int(float64(e[2]) / (1.0 + math.Sqrt(float64(d2))))
				}
			}
			if x < sum {
				o, x = []int{i, j}, sum
			}
		}
	}
	return o
}
