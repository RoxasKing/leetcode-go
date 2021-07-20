package main

import (
	"math"
)

func addRungs(rungs []int, dist int) int {
	n := len(rungs)
	out := int(math.Ceil(float64(rungs[0])/float64(dist))) - 1
	for i := 0; i < n-1; i++ {
		if rungs[i]+dist < rungs[i+1] {
			out += int(math.Ceil(float64(rungs[i+1]-rungs[i])/float64(dist))) - 1
		}
	}
	return out
}
