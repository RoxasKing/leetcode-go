package main

import "math/rand"

// Difficulty:
// Medium

// Tags:
// Radomized

type Solution struct {
	r, x, y float64
}

func Constructor(radius float64, x_center float64, y_center float64) Solution {
	return Solution{radius, x_center, y_center}
}

func (this *Solution) RandPoint() []float64 {
	x, y := 1.0, 1.0
	for x*x+y*y >= 1 {
		x = rand.Float64()*2 - 1
		y = rand.Float64()*2 - 1
	}
	return []float64{this.x + this.r*x, this.y + this.r*y}
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(radius, x_center, y_center);
 * param_1 := obj.RandPoint();
 */
